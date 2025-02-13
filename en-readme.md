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
-   [Redis](https://github.com/redis/go-redis)

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
APP_URL=127.0.0.1

DB_DATABASE=go_web_http
DB_USERNAME=your-username
DB_PASSWORD=your-password
MONGODB_DATABASE=goWebHttp
MONGODB_USERNAME=your-username
MONGODB_PASSWORD=your-password
REDIS_USERNAME=your-username
REDIS_PASSWORD=your-password
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
./bin/migrate create -ext sql -dir migrations -seq create_master_satker
```

4. Launch the app
```bash
go run .
# OR
air # Run with Air
```

<h3 id="develop-with-docker">🐳 Develop w/ Docker</h3>

-   Clone the repository:

```bash
git clone https://github.com/alfianchii/go-web-http
cd go-web-http
```

-   Copy `.env.example` file with `cp .env.example .env` and configure database:

```conf
APP_URL=0.0.0.0

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

-   Make sure you have Docker Compose installed and run:

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

<h2 id="production">🌐 Production</h2>

<h3 id="deployment-docker-vps">🐳 Deployment w/ Docker (use Virtual Private Server)</h3>

-   Clone the repository w/ SSH method `git clone git@github.com:alfianchii/go-web-http` and go to the directory with `cd go-web-http` command.

-   Copy `.env.example` file to `.env` and do configs.

```conf
APP_URL=127.0.0.1

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

-   Make sure on your VPS you have Docker Compose installed and run:

```bash
docker compose -f ./docker-compose.prod.yaml up -d --build
```

- Setup your domain and SSL certificate with Nginx configuration:

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

- Setup SSL certificate with Certbot:

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

<h2 id="support">💌 Support me</h2>

<p>You can support me on the Trakteer platform! Your support will be very meaningful. Like, just giving a star to this project is already greatly appreciated~!</p>

<a href="https://trakteer.id/alfianchii/tip" target="_blank"><img id="wse-buttons-preview" src="https://cdn.trakteer.id/images/embed/trbtn-red-5.png" height="40" style="border:0px;height:40px;" alt="Trakteer Me"></a>

<h2 id="contribution">🤝 Contributing</h2>

<p>Contributions, issues, and feature requests are highly appreciated as this application is far from perfect. Please do not hesitate to make a pull request and make changes to this project!</p>

<h2 id="license">📝 License</h2>

go-web-http is open-sourced software licensed under the [MIT License](./LICENSE).

<h2 id="author">🧍 Author</h2>

<p>go-web-http is created by <a href="https://instagram.com/alfianchii">Alfian</a>.</p>
