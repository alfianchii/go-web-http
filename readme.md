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

-   http://localhost:3333/
- - GET (Websocket)
-   http://localhost:3333/login
- - POST
- - - Param(s): username, password
-   http://localhost:3333/logout
- - POST
- - - Header(s): Cookie
-   http://localhost:3333/about
- - GET
- - - Header(s): Cookie
- - POST
- - - Param(s): email, subject, content
-   http://localhost:3333/greet/{name}
- - GET
-   http://localhost:3333/search?q={query}
- - GET
- - - Param(s): q

<h2 id="installation">💻 Installation</h2>

<h3 id="develop-yourself">🏃‍♂️ Develop by yourself</h3>

1. Clone repository

```bash
git clone https://github.com/alfianchii/go-web-http
cd go-web-http
```

2. Database configuration through the `./config/app.go` file
```go
const (
	Port = 3333
	Host = "localhost"
	DBHost = "127.0.0.1"
	DBPort = 5432 // This is the default port for PostgreSQL
	DBUser = "my-username"
	DBName = "my-database"
)
```

3. Launch the app
```bash
go mod tidy
go mod verify
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

web-http-go is open-sourced software licensed under the [MIT License](./LICENSE).

<h2 id="author">🧍 Author</h2>

<p>web-http-go is created by <a href="https://instagram.com/alfianchii">Alfian</a>.</p>
