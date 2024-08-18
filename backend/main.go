package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    InitDB()

    r.POST("/settings", SaveSettings)
    r.POST("/ipfs/upload", UploadToIPFS)
    r.GET("/ipfs/fetch/:cid", FetchFromIPFS)
    r.POST("/clickhouse/ingest", IngestIntoClickHouse)
    r.GET("/clickhouse/query", QueryClickHouse)

    r.Run(":8080")
}
