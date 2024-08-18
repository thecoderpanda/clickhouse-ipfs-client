package main

import (
    "database/sql"
    _ "github.com/ClickHouse/clickhouse-go/v2"
)

// Initialize and manage ClickHouse connection and operations here.

// Example function to query data from ClickHouse.
func queryClickHouse(query string) (*sql.Rows, error) {
    dsn := "tcp://" + config.ClickHouseURL + "?username=" + config.ClickHouseUser + "&password=" + config.ClickHousePass
    db, err := sql.Open("clickhouse", dsn)
    if err != nil {
        return nil, err
    }
    defer db.Close()

    return db.Query(query)
}
