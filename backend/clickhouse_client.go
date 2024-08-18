package main

import (
    "database/sql"
    "fmt"
    _ "github.com/ClickHouse/clickhouse-go"
)

func ingestDataIntoClickHouse(cid string, data string) error {
    settings := getSettings()
    dsn := fmt.Sprintf("tcp://%s?username=%s&password=%s&database=default",
        settings.ClickHouseURL, settings.ClickHouseUser, settings.ClickHousePass)

    conn, err := sql.Open("clickhouse", dsn)
    if err != nil {
        return err
    }
    defer conn.Close()

    _, err = conn.Exec("INSERT INTO ipfs_data (cid, data) VALUES (?, ?)", cid, data)
    return err
}

func queryDataFromClickHouse(cid string) (string, error) {
    settings := getSettings()
    dsn := fmt.Sprintf("tcp://%s?username=%s&password=%s&database=default",
        settings.ClickHouseURL, settings.ClickHouseUser, settings.ClickHousePass)

    conn, err := sql.Open("clickhouse", dsn)
    if err != nil {
        return "", err
    }
    defer conn.Close()

    var data string
    err = conn.QueryRow("SELECT data FROM ipfs_data WHERE cid = ?", cid).Scan(&data)
    return data, err
}

func getSettings() Settings {
    var settings Settings
    db.First(&settings)
    return settings
}
