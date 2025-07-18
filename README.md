# 🚀 GoFr Data Loader Service

A microservice built using [GoFr](https://github.com/tkrajina/gofr) to support file uploads and preview functionality for data ingestion workflows.

---

## ✨ Features

- `/upload` endpoint
- Upload `.csv` and `.json` files
- Previews first 5 rows/items of the uploaded file
- Clean error handling for unsupported or missing files

---

## 🔧 Tech Stack

- Language: Go
- Framework: [GoFr](https://github.com/tkrajina/gofr)
- API: RESTful

---

## 📦 How to Run

```bash
go run main.go
