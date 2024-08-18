
This project is a Go-based web application with a frontend and backend that facilitates integration between IPFS (InterPlanetary File System) and ClickHouse. The application allows users to upload files to IPFS, fetch data using a CID, and ingest that data into ClickHouse. Additionally, users can view the ingested data directly from the dashboard.

## Project Structure

```plaintext
go-ipfs-clickhouse-app/
│
├── backend/
│   ├── main.go
│   ├── handlers.go
│   ├── models.go
│   ├── db.go
│   ├── clickhouse_client.go
│   ├── ipfs_client.go
│   └── go.mod
│
├── frontend/
│   ├── public/
│   ├── src/
│   │   ├── App.vue
│   │   ├── main.js
│   │   ├── components/
│   │   │   ├── Settings.vue
│   │   │   ├── Dashboard.vue
│   │   │   └── DataTable.vue
│   ├── package.json
│   └── vue.config.js
└── README.md
```

## Backend Setup

### Prerequisites

- Go 1.16+
- ClickHouse server running
- IPFS node (optional, can use public nodes)

### Steps

1. **Navigate to the `backend/` directory**:
    ```bash
    cd backend/
    ```

2. **Initialize the Go module**:
    ```bash
    go mod init go-ipfs-clickhouse-app
    ```

3. **Install dependencies**:
    ```bash
    go get github.com/gin-gonic/gin gorm.io/gorm gorm.io/driver/sqlite github.com/ClickHouse/clickhouse-go
    ```

4. **Run the backend server**:
    ```bash
    go run main.go
    ```

    The backend server will run on `http://localhost:8080`.

## Frontend Setup

### Prerequisites

- Node.js 14+
- npm or yarn

### Steps

1. **Navigate to the `frontend/` directory**:
    ```bash
    cd frontend/
    ```

2. **Install dependencies**:
    ```bash
    npm install
    ```

3. **Run the frontend server**:
    ```bash
    npm run serve
    ```

    The frontend server will run on `http://localhost:8080`.

## Operation

### Settings Page

- Access the settings page from the frontend UI.
- Configure the IPFS Node URL, ClickHouse URL, and ClickHouse credentials (username and password).
- Click on "Save Settings" to store the configuration.

### Dashboard Page

- **Upload to IPFS**:
  - Select a file from your local machine.
  - Click on "Upload" to upload the file to the configured IPFS node. The CID will be displayed upon successful upload.
  
- **Fetch from IPFS**:
  - Enter a CID to fetch data from IPFS.
  - The data will be fetched, ingested into ClickHouse, and displayed on the dashboard.

- **View Data**:
  - The data ingested into ClickHouse will be displayed in a table format, allowing users to view the results of their operations.

## Local Storage

- **CIDs and User Credentials**:
  - Stored locally in an SQLite database (`app.db`) within the backend directory.
  - The settings and CID data are persisted and can be retrieved or updated as needed.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

## Contributing

Contributions are welcome! Please fork this repository and submit a pull request for any improvements or bug fixes.

## Acknowledgements

- [Gin Web Framework](https://gin-gonic.com/)
- [GORM ORM Library](https://gorm.io/)
- [ClickHouse](https://clickhouse.com/)
- [IPFS](https://ipfs.io/)


