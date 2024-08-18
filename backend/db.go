package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
    var err error
    db, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&Settings{}, &CIDData{})
}

func GetNodeURL() string {
    var settings Settings
    db.First(&settings)
    return settings.IPFSNodeURL
}
