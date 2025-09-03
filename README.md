# skelgo

A CLI tool for quickly scaffolding Go REST API projects with [Gin](https://github.com/gin-gonic/gin), [GORM](https://gorm.io/), and PostgreSQL.  
skelgo generates a modular folder structure, starter files, and boilerplate code to help you get started building production-ready Go web APIs fast.

---

## ✨ Features

- Generates a clean, idiomatic Go project structure
- Integrates [Gin](https://github.com/gin-gonic/gin) for HTTP routing
- Uses [GORM](https://gorm.io/) for ORM and PostgreSQL support
- Includes starter code for user authentication, middleware, and more
- Provides migration documentation and helpers using [golang-migrate](https://github.com/golang-migrate/migrate)
- Manual dependency wiring for easy extensibility

---

## 🛠️ Requirements

- [Go](https://golang.org/dl/) 1.18 or newer
- [golang-migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) (for database migrations)
- [Git](https://git-scm.com/) (for initializing repositories, optional)
- PostgreSQL database (for running the generated project)

---

## 🚀 Installation

Clone this repository and build the CLI tool:

```sh
git clone https://github.com/JerryJeager/skelgo.git
cd skelgo
go build -o skelgo
```

Or install directly with Go:

```sh
go install github.com/JerryJeager/skelgo@latest
```

Make sure the resulting `skelgo` binary is in your `PATH`.

---

## ⚡ Usage

To scaffold a new project, run:

```sh
skelgo init <project-name> <module-path>
```

- `<project-name>`: The name of your new project folder (e.g., `myapp`)
- `<module-path>`: Your Go module path (e.g., `github.com/yourusername/myapp`)

Example:

```sh
skelgo init myapp github.com/yourusername/myapp
```

This will create a new folder `myapp` with all the starter files and folders.

---

## 📝 After Generation

1. **Install dependencies:**
   ```sh
   cd myapp
   go mod tidy
   ```

2. **Set up your environment:**
   - Copy `.env.example` to `.env` and fill in your database credentials and other settings.

3. **Install golang-migrate CLI:**
   - [Installation instructions](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation)
   - Required for running and creating database migrations.

4. **Run database migrations:**
   ```sh
   migrate -path internal/db/migrations -database "postgresql://postgres:YOURPASSWORD@localhost:5432/YOURDB?sslmode=disable" -verbose up
   ```

5. **Run your server:**
   ```sh
   go run main.go
   ```

---

## 📚 Documentation

- See the generated [`docs/migrations.md`](docs/migrations.md) in your project for migration usage.
- Review the generated `README.md` in your new project for more details.

---

## 🤝 Contributing

Pull requests and issues are welcome! Please open an issue to discuss your ideas or report bugs.

---

## License

MIT
