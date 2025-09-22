# Golang with Echo – Starter Project

Project ini adalah boilerplate/service starter berbasis **Golang + Echo** dengan integrasi ke **PostgreSQL, Redis, dan MinIO**.  
Cocok untuk microservice atau REST API modern dengan dependency yang umum dipakai.

---

## 🛠 Tech Stack
- **Golang**: `go1.23.2 windows/amd64`
- **Echo**: web framework ringan & cepat
- **GORM**: ORM untuk PostgreSQL
- **PostgreSQL**: database relasional
- **Redis**: in-memory cache & message broker
- **MinIO**: object storage (S3 compatible)
- **Docker**: untuk container service
- **lumberjack**: log rotation per hari

---

## 📂 Struktur Project
```text
.
├── config/         # konfigurasi (DB, Redis, MinIO, Logger)
├── internal/
│   ├── handler/    # HTTP handler (Echo)
│   ├── model/      # Model GORM
│   ├── repository/ # Repository (DB access)
│   ├── service/    # Business logic service
│   └── usecase/    # Use case orchestration
│   └── migration/  # AutoMigrate DB
│   └── routes/     # Routing Echo
├── main.go         # Entry point
└── go.mod
