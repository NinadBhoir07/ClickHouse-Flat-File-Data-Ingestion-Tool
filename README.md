# ClickHouse & Flat File Data Ingestion Tool

A web-based application that facilitates **bidirectional data ingestion** between a **ClickHouse database** and **Flat Files (CSV)** with column selection, JWT-based authentication, and real-time ingestion feedback.

---

## 🚀 Features

- 🔄 **Bidirectional Ingestion**: ClickHouse → Flat File, Flat File → ClickHouse
- 🔐 **JWT Token Authentication** for ClickHouse
- 📄 **Flat File Support** with delimiter configuration
- 📑 **Schema Discovery** and column selection
- 📊 **Multi-table JOIN support** for ClickHouse ingestion
- 🧪 **Data Preview** before ingestion
- 📈 **Ingestion Progress Bar**
- 📥 **Completion Summary** with record count

---

## 🧰 Tech Stack

- **Backend**: Go (Gin Web Framework)
- **Frontend**: HTML, CSS, JavaScript
- **ClickHouse Client**: Configurable with host, port, and JWT

---

## 📂 Project Structure

clickhouse_flatfile_ingestion_tool/
├── backend/
│   └── main.go                   # Go backend with Gin, handles ingestion logic and multi-table JOIN
│
├── frontend/
│   ├── index.html               # Main UI for interaction with the tool
│   └── style.css                # Basic styling for the frontend
│
├── sample_data.csv              # Sample CSV file for testing Flat File ingestion
├── README.md                    # Project documentation
