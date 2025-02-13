<div align="center">
	<h1>go-web-httpにようこそ！! 👋</h1>
	<a href="https://github.com/alfianchii/go-web-http/blob/main/en-readme.md" target="_blank">EN Readme!</a>
</div>

[![全部Contributors](https://img.shields.io/github/contributors/alfianchii/go-web-http)](https://github.com/alfianchii/go-web-http/graphs/contributors)
![GitHubの最後のコミット](https://img.shields.io/github/last-commit/alfianchii/go-web-http)

---

<h2 id="about">🤔 go-web-httpって何？</h2>

<p>go-web-httpは、Goプログラミング言語を使って作られたシンプルなWeb APIなんだ。このAPIは、Goを使ってWeb APIを作る方法を学ぶために作られたんだよね。</p>

<h2 id="features">🤨 go-web-http中には、どんなfeatureがあるの？</h2>

-   [Chi Framework](https://github.com/go-chi/chi)
-   [Go Postgres Driver](https://github.com/lib/pq)
-   [sqlx](https://github.com/jmoiron/sqlx)
-   [websocket](https://github.com/gorilla/websocket)
-   [golang-migrate](https://github.com/golang-migrate/migrate)
-   [godotenv](https://github.com/joho/godotenv)
-   [mongodb](https://github.com/mongodb/mongo-go-driver)
-   [Auth w/ JWT](https://github.com/golang-jwt/jwt)
-   [Redis](https://github.com/redis/go-redis)

<h2 id="routes">👤 ルート</h2>

### 👨‍🏫 アドミン

-   http://localhost:3333/admin
- - GET
-   http://localhost:3333/admin/dashboard
- - GET
-   http://localhost:3333/admin/settings
- - GET
-   http://localhost:3333/admin/books/{title}/page/{page}
- - GET
-   http://localhost:3333/admin/satker
- - GET

### 🧗 ユーザー

-   http://localhost:3333/register
- - GET
-   http://localhost:3333/user
- - POST (API)
- - Param(s)
- - - username
- - - password
- - - email
-   http://localhost:3333/
- - GET (Chat w/ WebSocket)
- - Header(s)
- - - Cookie
-   http://localhost:3333/login
- - POST (API)
- - Param(s)
- - - username
- - - password
-   http://localhost:3333/logout
- - POST (API)
- - Header(s)
- - - Cookie
- - - Authorization: Bearer {token}
-   http://localhost:3333/validate-jwt
- - POST (API)
- - Header(s)
- - - Cookie
- - - Authorization: Bearer {token}
-   http://localhost:3333/about
- - GET
- - Header(s)
- - - Cookie
- - POST (API)
- - Param(s)
- - - email
- - - subject
- - - content
-   http://localhost:3333/greet/{name}
- - GET (API)
-   http://localhost:3333/search?q={query}
- - GET (API)
- - Param(s)
- - - q

<h2 id="installation">💻 設置</h2>

<h3 id="develop-yourself">🏃‍♂️ 自分で開発してみて</h3>

1. Repositoryをクローンして、dependenciesをインストールする
```bash
git clone https://github.com/alfianchii/go-web-http
cd go-web-http
go mod tidy
go mod verify
cp .env.example .env
```

2. `.env`ファイルを通じてdatabaseを設定する
```bash
DB_DATABASE=go_web_http
DB_USERNAME=your-username
DB_PASSWORD=your-password
MONGODB_DATABASE=goWebHttp
MONGODB_USERNAME=your-username
MONGODB_PASSWORD=your-password
REDIS_USERNAME=your-username
REDIS_PASSWORD=your-password
```

3. ローカルに、[golang-migrate](https://github.com/golang-migrate/migrate)をインストールして、migrationを実行する
```bash
GOBIN=$(pwd)/bin go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
./bin/migrate -database "postgres://your-username:your-password@127.0.0.1:5432/go_web_http?sslmode=disable" -path ./migrations up
# Migrationをダウングレードする
./bin/migrate -database "postgres://your-username:your-password@127.0.0.1:5432/go_web_http?sslmode=disable" -path ./migrations down
```

- さらに、自分でmigrationsを作成したい場合は、次のcommandを使用できます：
```bash
./bin/migrate create -ext sql -dir migrations -seq create_<table_name>
```
- 例えば：
```bash
./bin/migrate create -ext sql -dir migrations -seq create_master_satker
```

4. Appを起動する
```bash
go run .
# OR
air # Airと実行する
```

<h3 id="develop-with-docker">🐳 Dockerで開発する</h3>

-   Repositoryをクローンする:

```bash
git clone https://github.com/alfianchii/go-web-http
cd go-web-http
co .env.example .env
```

-  `.env`ファイルを通じてcore depsを設定する：

```conf
APP_URL=127.0.0.1

DB_HOST=postgres
DB_DATABASE=go_web_http
DB_USERNAME=gowebhttp
DB_PASSWORD=gowebhttp1@1

MONGODB_HOST=mongo
MONGODB_DATABASE=goWebHttp
MONGODB_USERNAME=gowebhttp
MONGODB_PASSWORD=gowebhttp1@1

REDIS_HOST=redis
REDIS_USERNAME=gowebhttp
REDIS_PASSWORD=gowebhttp1@1
```

-   Docker Composeがインストールされてるか確認してね。それから、このcommandを実行してみて：

```bash
docker compose up --build -d
```

-   Pages
-   -   App: `http://0.0.0.0:3333`

<h4 id="docker-commands">🔐 Commands</h4>

-   Go
-   -   `docker compose exec app sh`
-   -   `docker compose exec app go mod tidy`
-   -   `docker compose exec app go get <deps>`
-   -   `docker compose exec app go install <deps>`
-   -   `docker compose exec app go build -o main .`
-   -   Etc

-   Redis
-   -   `docker compose exec redis bash`
-   -   `docker compose exec redis redis-cli`
-   -   Etc

-   Mongo
-   -   `docker compose exec mongo bash`
-   -   `docker compose exec mongo mongosh -u gowebhttp -p gowebhttp1@1`
-   -   Etc

-   Postgres
-   -   `docker compose exec postgres bash`
-   -   `docker compose exec postgres psql -U gowebhttp -d go_web_http`
-   -   Etc

<h2 id="production">🌐 本番環境</h2>

<h3 id="deployment-docker-vps">🐳 Dockerを使ってデプロイするよ (Virtual Private Serverを使う感じで!)</h3>

-   リポジトリをSSHでクローンしてみて！`git clone git@github.com:alfianchii/go-web-http`って打てばOKだよ。それから、`cd go-web-http`でディレクトリに移動してね！

-   `.env.example`ファイルをコピーして`.env`にしてね！それから設定をいじってみて：

```conf
APP_URL=0.0.0.0

DB_HOST=postgres
DB_DATABASE=go_web_http
DB_USERNAME=your-vps-username
DB_PASSWORD=your-vps-password

MONGODB_HOST=mongo
MONGODB_DATABASE=goWebHttp
MONGODB_USERNAME=your-vps-username
MONGODB_PASSWORD=your-vps-password

REDIS_HOST=redis
REDIS_USERNAME=your-vps-username
REDIS_PASSWORD=your-vps-password
```

-   VPSにDocker Composeがインストールされてることを確認してね。それから、このcommandを実行してみて：

```bash
docker compose -f ./docker-compose.prod.yaml up -d --build
```

- ドメインとSSL証明書を設定して、Nginxの設定をするよ：

```nginx
server {
  server_name your-domain.com www.your-domain.com;

  location / {
    proxy_pass http://127.0.0.1:3333;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
  }

  location /chats {
    proxy_pass http://127.0.0.1:3333;
    proxy_http_version 1.1;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection "upgrade";
    proxy_set_header Host $host;
  }

  error_log /var/log/nginx/your-domain_error.log;
  access_log /var/log/nginx/your-domain_access.log;
}
```

- Certbotを使ってSSL証明書を設定しよう！コマンドはこれ：

```bash
sudo apt install certbot python3-certbot-nginx
sudo certbot --nginx -d your-domain.com -d www.your-domain.com
sudo ln -s /etc/nginx/sites-available/your-domain.com /etc/nginx/sites-enabled/
sudo systemctl reload nginx
```

<h4 id="docker-commands">🔐 Commands</h4>

-   Go
-   -   `docker compose -f ./docker-compose.prod.yaml exec app sh`
-   -   `docker compose -f ./docker-compose.prod.yaml exec app go mod tidy`
-   -   `docker compose -f ./docker-compose.prod.yaml exec app go get <deps>`
-   -   `docker compose -f ./docker-compose.prod.yaml exec app go install <deps>`
-   -   `docker compose -f ./docker-compose.prod.yaml exec app go build -o main .`
-   -   Etc

-   Redis
-   -   `docker compose -f ./docker-compose.prod.yaml exec redis bash`
-   -   `docker compose -f ./docker-compose.prod.yaml exec redis redis-cli`
-   -   Etc

-   Mongo
-   -   `docker compose -f ./docker-compose.prod.yaml exec mongo bash`
-   -   `docker compose -f ./docker-compose.prod.yaml exec mongo mongosh -u gowebhttp -p gowebhttp1@1`
-   -   Etc

-   Postgres
-   -   `docker compose -f ./docker-compose.prod.yaml exec postgres bash`
-   -   `docker compose -f ./docker-compose.prod.yaml exec postgres psql -U gowebhttp -d go_web_http`
-   -   Etc

<h2 id="support">💌 応援してね</h2>

<p>Trakteerプラットフォームで私を応援してくれると嬉しいです！みんなのスポートはとても意味のあるものになります。例えば、このプロジェクトに星をつけるだけでも大変感謝しています〜！</p>

<a href="https://trakteer.id/alfianchii/tip" target="_blank"><img id="wse-buttons-preview" src="https://cdn.trakteer.id/images/embed/trbtn-red-5.png" height="40" style="border:0px;height:40px;" alt="Trakteer Me"></a>

<h2 id="contribution">🤝 貢献する</h2>

<p>このappはまだまだ完璧ではないため、contributions、issues、feature requestsは大変ありがたいです。PRをためらわずに作成して、このプロジェクトに変更を加えてください！</p>

<h2 id="license">📝 ライセンス</h2>

go-web-http is open-sourced software licensed under the [MIT License](./LICENSE).

<h2 id="author">🧍 作成者</h2>

<p>go-web-httpは<a href="https://instagram.com/alfianchii">Alfian</a>によって作られました</a>.</p>
