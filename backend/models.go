package main

import (
    "gorm.io/gorm"
)

type Settings struct {
    gorm.Model
    IPFSNodeURL     string
    ClickHouseURL   string
    ClickHouseUser  string
    ClickHousePass  string
}

type CIDData struct {
    gorm.Model
    CID  string
    Data string
}
