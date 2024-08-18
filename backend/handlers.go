package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func SaveSettings(c *gin.Context) {
    var settings Settings
    if err := c.ShouldBindJSON(&settings); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.Save(&settings)
    c.JSON(http.StatusOK, gin.H{"status": "settings saved"})
}

func UploadToIPFS(c *gin.Context) {
    file, _ := c.FormFile("file")
    nodeURL := GetNodeURL()
    cid, err := uploadFileToIPFS(file, nodeURL)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    db.Save(&CIDData{CID: cid})
    c.JSON(http.StatusOK, gin.H{"cid": cid})
}

func FetchFromIPFS(c *gin.Context) {
    cid := c.Param("cid")
    nodeURL := GetNodeURL()
    data, err := fetchFromIPFS(cid, nodeURL)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": data})
}

func IngestIntoClickHouse(c *gin.Context) {
    var ingestRequest struct {
        CID  string `json:"cid"`
        Data string `json:"data"`
    }
    if err := c.ShouldBindJSON(&ingestRequest); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    err := ingestDataIntoClickHouse(ingestRequest.CID, ingestRequest.Data)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "data ingested"})
}

func QueryClickHouse(c *gin.Context) {
    cid := c.Query("cid")
    data, err := queryDataFromClickHouse(cid)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": data})
}
