<h1 align="center">Welcome to go-web-http! 👋</h1>

[![All Contributors](https://img.shields.io/github/contributors/alfianchii/go-web-http)](https://github.com/alfianchii/go-web-http/graphs/contributors)
![GitHub last commit](https://img.shields.io/github/last-commit/alfianchii/go-web-http)

---

<h2 id="about">🤔 What is go-web-http?</h2>

<p>go-web-http is a simple web APIs that are created using the Go programming language. This app is created to learn how to create a web APIs using the Go programming language.</p>

<h2 id="features">🤨 What features are available in go-web-http?</h2>

-   [Chi Framework](https://github.com/go-chi/chi)
-   [Go Postgres Driver](https://github.com/lib/pq)
-   [sqlx](https://github.com/jmoiron/sqlx)
-   [websocket](https://github.com/gorilla/websocket)
-   [golang-migrate](https://github.com/golang-migrate/migrate)
-   [godotenv](https://github.com/joho/godotenv)
-   [mongodb](https://github.com/mongodb/mongo-go-driver)
-   [Auth w/ JWT](https://github.com/golang-jwt/jwt)

<h2 id="routes">👤 Routes</h2>

### 👨‍🏫 Admin

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

### 🧗 User

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

<h2 id="installation">💻 Installation</h2>

<h3 id="develop-yourself">🏃‍♂️ Develop by yourself</h3>

1. Clone repository and install dependencies
```bash
git clone https://github.com/alfianchii/go-web-http
cd go-web-http
go mod tidy
go mod verify
cp .env.example .env
```

2. Database configuration through the `.env` file
```bash
DB_DATABASE=go_web_http
DB_USERNAME=your-username
DB_PASSWORD=your-password
MONGODB_DATABASE=goWebHttp
MONGODB_USERNAME=your-username
MONGODB_PASSWORD=your-password
```

3. Install [golang-migrate](https://github.com/golang-migrate/migrate) locally and run the migration
```bash
GOBIN=$(pwd)/bin go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
./bin/migrate -database "postgres://your-username:your-password@127.0.0.1:5432/go_web_http?sslmode=disable" -path ./migrations up
# Down the migration
./bin/migrate -database "postgres://your-username:your-password@127.0.0.1:5432/go_web_http?sslmode=disable" -path ./migrations down
```

- Additionally, if you want make your own migrations, you can use the following command:
```bash
./bin/migrate create -ext sql -dir migrations -seq create_<table_name>
```
- For example:
```bash
./bin/migrate create -ext sql -dir migrations -seq create_mst_satker
```

4. Launch the app
```bash
go run .
# OR
air # Run with Air
```

<h2 id="support">💌 Support me</h2>

<p>You can support me on the Trakteer platform! Your support will be very meaningful. Like, just giving a star to this project is already greatly appreciated~!</p>

<a href="https://trakteer.id/alfianchii/tip" target="_blank"><img id="wse-buttons-preview" src="https://cdn.trakteer.id/images/embed/trbtn-red-5.png" height="40" style="border:0px;height:40px;" alt="Trakteer Me"></a>

<h2 id="contribution">🤝 Contributing</h2>

<p>Contributions, issues, and feature requests are highly appreciated as this application is far from perfect. Please do not hesitate to make a pull request and make changes to this project!</p>

<h2 id="license">📝 License</h2>

go-web-http is open-sourced software licensed under the [MIT License](./LICENSE).

<h2 id="author">🧍 Author</h2>

<p>go-web-http is created by <a href="https://instagram.com/alfianchii">Alfian</a>.</p>
