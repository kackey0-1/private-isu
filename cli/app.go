package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

const (
	postsPerPage  = 20
	ISO8601Format = "2006-01-02T15:04:05-07:00"
	UploadLimit   = 10 * 1024 * 1024 // 10MB
)

type User struct {
	ID          int       `db:"id"`
	AccountName string    `db:"account_name"`
	Passhash    string    `db:"passhash"`
	Authority   int       `db:"authority"`
	DelFlg      int       `db:"del_flg"`
	CreatedAt   time.Time `db:"created_at"`
}

type Post struct {
	ID           int       `db:"id"`
	UserID       int       `db:"user_id"`
	Imgdata      []byte    `db:"imgdata"`
	Body         string    `db:"body"`
	Mime         string    `db:"mime"`
	CreatedAt    time.Time `db:"created_at"`
	CommentCount int
	Comments     []Comment
	User         User
	CSRFToken    string
}

type Comment struct {
	ID        int       `db:"id"`
	PostID    int       `db:"post_id"`
	UserID    int       `db:"user_id"`
	Comment   string    `db:"comment"`
	CreatedAt time.Time `db:"created_at"`
	User      User
}

// mimeToExt: MIMEタイプから拡張子を決定する関数
func mimeToExt(mime string) (string, error) {
	switch mime {
	case "image/jpeg":
		return "jpg", nil
	case "image/png":
		return "png", nil
	case "image/gif":
		return "gif", nil
	default:
		return "", fmt.Errorf("unsupported mime type: %s", mime)
	}
}

// getImagesCLI: 100件ずつバッチで画像を取得し、保存する処理をCLIとして実行
func getImagesCLI() {
	batchSize := 100       // 1回のクエリで取得する件数
	totalRecords := 100000 // 処理する総レコード数
	totalBatches := totalRecords / batchSize

	for batch := 0; batch < totalBatches; batch++ {
		offset := batch * batchSize

		// 100件分のデータを取得
		posts := []Post{}
		query := "SELECT * FROM `posts` ORDER BY `id` LIMIT ? OFFSET ?"
		err := db.Select(&posts, query, batchSize, offset)
		if err != nil {
			log.Print(err)
			return
		}

		// 取得したポストデータを処理
		for _, post := range posts {
			// MIMEタイプから拡張子を取得
			ext, err := mimeToExt(post.Mime)
			if err != nil {
				log.Print(err)
				continue
			}

			// 画像データをファイルに書き出し
			filePath := fmt.Sprintf("../webapp/public/image/%d.%s", post.ID, ext)
			err = writeImageToFile(filePath, post.Imgdata)
			if err != nil {
				log.Print(err)
				continue
			}
		}
		log.Printf("Processed batch %d/%d (offset: %d)", batch+1, totalBatches, offset)
	}
	log.Println("All posts processed.")
}

// 画像データをファイルに保存する関数
func writeImageToFile(filePath string, imgData []byte) error {
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	_, err = f.Write(imgData)
	if err != nil {
		return fmt.Errorf("failed to write image data: %w", err)
	}

	log.Printf("Image saved to %s", filePath)
	return nil
}

func main() {
	// DB接続設定（環境変数から読み込み）
	host := os.Getenv("ISUCONP_DB_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("ISUCONP_DB_PORT")
	if port == "" {
		port = "3306"
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Failed to read DB port number from an environment variable ISUCONP_DB_PORT.\nError: %s", err.Error())
	}
	user := os.Getenv("ISUCONP_DB_USER")
	if user == "" {
		user = "root"
	}
	password := os.Getenv("ISUCONP_DB_PASSWORD")
	dbname := os.Getenv("ISUCONP_DB_NAME")
	if dbname == "" {
		dbname = "isuconp"
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %s.", err.Error())
	}
	defer db.Close()

	// 画像をCLI経由で処理
	getImagesCLI()
}
