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
- - - Authorization: Bearer {token}
-   http://localhost:3333/login
- - POST (API)
- - Param(s)
- - - username
- - - password
-   http://localhost:3333/logout
- - POST (API)
- - Header(s)
- - - Cookie
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

2. `.env`ファイルを通じてDatabaseを設定する
```bash
DB_DATABASE=go_web_http
DB_USERNAME=your-username
DB_PASSWORD=your-password
MONGODB_DATABASE=goWebHttp
MONGODB_USERNAME=your-username
MONGODB_PASSWORD=your-password
```

3. ローカルに、[golang-migrate](https://github.com/golang-migrate/migrate)をインストールして、migrationを実行する
```bash
GOBIN=$(pwd)/bin go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
./bin/migrate -database "postgres://your-username:your-password@127.0.0.1:5432/go_web_http?sslmode=disable" -path ./migrations up
# Migrationをダウングレードする
./bin/migrate -database "postgres://your-username:your-password@127.0.0.1:5432/go_web_http?sslmode=disable" -path ./migrations down
```

- さらに、自分でmigrationsを作成したい場合は、次のcommandを使用できます:
```bash
./bin/migrate create -ext sql -dir migrations -seq create_<table_name>
```
- 例えば：
```bash
./bin/migrate create -ext sql -dir migrations -seq create_mst_satker
```

4. Appを起動する
```bash
go run .
# OR
air # Airと実行する
```

<h2 id="support">💌 応援してね</h2>

<p>Trakteerプラットフォームで私を応援してくれると嬉しいです！みんなのスポートはとても意味のあるものになります。例えば、このプロジェクトに星をつけるだけでも大変感謝しています〜！</p>

<a href="https://trakteer.id/alfianchii/tip" target="_blank"><img id="wse-buttons-preview" src="https://cdn.trakteer.id/images/embed/trbtn-red-5.png" height="40" style="border:0px;height:40px;" alt="Trakteer Me"></a>

<h2 id="contribution">🤝 貢献する</h2>

<p>このappはまだまだ完璧ではないため、contributions、issues、feature requestsは大変ありがたいです。PRをためらわずに作成して、このプロジェクトに変更を加えてください！</p>

<h2 id="license">📝 ライセンス</h2>

go-web-http is open-sourced software licensed under the [MIT License](./LICENSE).

<h2 id="author">🧍 作成者</h2>

<p>go-web-httpは<a href="https://instagram.com/alfianchii">Alfian</a>によって作られました</a>.</p>
