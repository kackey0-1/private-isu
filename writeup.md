# 初回のやってみた

## HOW TO

- [remote deploy方法](https://github.com/Gurrium/private-isu/blob/main/webapp/scripts/remote/deploy)
  ```bash
  cd /home/isucon/private-isu/webapp/go
  /home/isucon/.local/go/bin/go build -o app
  sudo systemctl restart isu-go
  ```

## Monitoring項目整理

- WebServer
  - サーバーの処理能力を全て使えているか
  - 並列度が期待する並列度になっているか
- Database Access
  - N+1問題を見つける
  - Slow Queryを特定するための設定を入れる
    - 打ち手一覧
      - indexが貼られているcolumnをwhere句に入れる
      - indexを新規で貼る
  - IO負荷の高いdatabaseの場合は、以下の点を確認すること
    - buffer poolの活用
    - バイナリログの出力方法(sync_binlog)
    - ログのフラッシュタイミングの調整

## AWS上に環境構築

```bash
# Deploy
aws cloudformation create-stack --stack-name private-isu-v3 \
  --template-body file://cf/private-isu.yaml \
  --parameters ParameterKey=KeyPairName,ParameterValue=key-pair ParameterKey=GitHubUsername,ParameterValue=kackey0-1


sudo systemctl stop isu-ruby
sudo systemctl disable isu-ruby
sudo systemctl start isu-go
sudo systemctl enable isu-go
```

## Benchmark実行

```bash
sudo su - isucon
/home/isucon/private_isu.git/benchmarker/bin/benchmarker -u /home/isucon/private_isu.git/benchmarker/userdata -t http://52.196.250.55
```

## AWS上のリソース削除

```bash
aws cloudformation delete-stack --stack-name private-isu-v3
```

## 実際に解いてみた

最初の問題
```
{"pass":true,"score":0,"success":196,"fail":64,"messages":["リクエストがタイムアウトしました (GET /)","リクエストがタイムアウトしました (GET /@amy)","リクエストがタイムアウトしました (GET /@deanne)","リクエストがタイムアウトしました (GET /@denise)","リクエストがタイムアウトしました (GET /@earlene)","リクエストがタイムアウトしました (GET /@esperanza)","リクエストがタイムアウトしました (GET /@magdalena)","リクエストがタイムアウトしました (GET /@maxine)","リクエストがタイムアウトしました (GET /@tracie)","リクエストがタイムアウトしました (GET /@verna)","リクエストがタイムアウトしました (GET /@vickie)","リクエストがタイムアウトしました (POST /login)","リクエストがタイムアウトしました (POST /register)"]}
```

`/etc/mysql/mysql.conf.d/mysqld.cnf` を編集
```cnf
slow_query_log		= 1
slow_query_log_file	= /var/log/mysql/mysql-slow.log
long_query_time = 0.5
```


-> 以下のindexを追加
```sql
ALTER TABLE comments ADD INDEX comments_post_idx (post_id);
```

### ログフォーマットを修正して調査

```conf
server {
	## Log Setting
	log_format json escape=json '{"time":"$time_iso8601",
    '"host":"$remote_addr",'
		'"port":$remote_port,'
		'"method":"$request_method",'
		'"uri":"$request_uri",'
		'"status":"$status",'
		'"body_bytes":$body_bytes_sent,'
		'"referer":"$http_referer",'
		'"ua":"$http_user_agent",'
		'"request_time":"$request_time",'
		'"response_time":"$upstream_response_time"}';
	access_log /var/log/nginx/access.log json;
  # ...
}
```

```bash
scp isucon@52.194.167.18:/var/log/nginx/access.log webapp/logs/access.log
alp json --sort sum -r -m "posts/[0-9]+,/@\w+,/image/\d+" -o count,method,uri,min,avg,max,sum < webapp/logs/access.log
```
