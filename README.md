# Server Generator
#### Author: Bocaletto Luca

**Server-Gen** is a multi-platform Go application that collects system metrics (IP, OS, CPU, memory, users) and emails a templated report on a configurable schedule. It also exposes a `/healthz` endpoint for basic health checks.

---

## 🚀 Features

- Cron-style scheduling (`@every 30m`, `0 8 * * *`, etc.)  
- Collects:  
  - Network interfaces & IPs  
  - Hostname & OS version  
  - CPU utilization  
  - Memory usage  
  - Active users  
- Sends HTML/text email via SMTP with TLS + retry  
- Configuration via YAML + environment‐variable overrides  
- Built-in HTTP health endpoint (`/healthz`)  
- Structured JSON+console logging (zerolog)  
- Docker multi-stage build  
- CI with GitHub Actions (matrix builds, tests, linting)

---

## 📁 Repository Layout

```
server-gen/
├── .github/
│   └── workflows/ci.yml
├── cmd/
│   └── server-gen/
│       └── main.go
├── configs/
│   └── config.yaml
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── mailer/
│   │   └── mailer.go
│   └── sysinfo/
│       └── sysinfo.go
├── Dockerfile
├── go.mod
└── README.md
```

---

## ⚙️ Requirements

- Go ≥ 1.21  
- Docker (optional)  
- SMTP server (StartTLS or SMTPS)

---

## 🔧 Installation

```bash
# Clone the repo
git clone https://github.com/bocaletto-luca/server-gen.git
cd server-gen

# Build binary
go build -o server-gen ./cmd/server-gen
```

---

## 🛠️ Configuration

Copy `configs/config.yaml` and fill in your settings. You can also override any value via environment variables.

```yaml
# configs/config.yaml
schedule: "@every 30m"

smtp:
  host: "${SMTP_HOST}"
  port: ${SMTP_PORT}
  username: "${SMTP_USER}"
  password: "${SMTP_PASS}"
  from: "${SMTP_FROM}"
  to:
    - "${SMTP_TO1}"
    - "${SMTP_TO2}"

modules:
  - ip
  - os
  - cpu
  - mem
  - users

http:
  addr: ":8080"
```

**Env vars**  
```bash
export SMTP_HOST=smtp.example.com
export SMTP_PORT=587
export SMTP_USER=alert@example.com
export SMTP_PASS=supersecret
export SMTP_FROM=alert@example.com
export SMTP_TO1=admin1@example.com
export SMTP_TO2=admin2@example.com
```

---

## ▶️ Usage

```bash
./server-gen --config configs/config.yaml
```

- The app starts an HTTP server on `http://localhost:8080/healthz`.  
- It schedules jobs per `schedule` in the config.  
- Each run gathers metrics and emails the report.

---

## 🐳 Docker

Build and run with Docker:

```bash
# Build image
docker build -t bocaletto-luca/server-gen .

# Run container
docker run -d \
  -v $(pwd)/configs/config.yaml:/app/configs/config.yaml:ro \
  -e SMTP_HOST -e SMTP_PORT -e SMTP_USER -e SMTP_PASS \
  -e SMTP_FROM -e SMTP_TO1 -e SMTP_TO2 \
  --name server-gen \
  bocaletto-luca/server-gen
```

---

## 🧪 Testing & CI

- **Unit tests**:  
  ```bash
  go test ./internal/config
  go test ./internal/sysinfo
  go test ./internal/mailer
  ```
- **Lint & vet**:  
  ```bash
  go vet ./...
  ```
- **GitHub Actions** runs tests & builds for Go 1.19–1.21.

---

## 🤝 Contributing

1. Fork the repository  
2. Create a feature branch (`git checkout -b feat/your-feature`)  
3. Commit your changes (`git commit -m "feat: ..."`); run tests  
4. Push to your branch (`git push origin feat/your-feature`)  
5. Open a Pull Request

---

## 📄 License

This project is licensed under the **MGPL License**.  
See [LICENSE](LICENSE) for details.
