初回のやってみた

```bash
cd webapp/sql
curl -L -O https://github.com/catatsuy/private-isu/releases/download/img/dump.sql.bz2
bunzip2 dump.sql.bz2
cd ..

vim docker-compose.yml
# ruby/ -> golang/
docker compose build app
docker compose up -d
```