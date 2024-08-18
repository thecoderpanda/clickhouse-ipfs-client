package main

import (
    "encoding/json"
    "io/ioutil"
    "sync"
)

type Config struct {
    IPFSNodeURL    string `json:"ipfs_node_url"`
    IPFSUsername   string `json:"ipfs_username"`
    IPFSPassword   string `json:"ipfs_password"`
    ClickHouseURL  string `json:"clickhouse_url"`
    ClickHouseUser string `json:"clickhouse_user"`
    ClickHousePass string `json:"clickhouse_pass"`
}

var (
    config      Config
    configMutex sync.RWMutex
)

// loadConfig reads the configuration from config.json.
func loadConfig() error {
    data, err := ioutil.ReadFile("config.json")
    if err != nil {
        return err
    }

    configMutex.Lock()
    defer configMutex.Unlock()

    return json.Unmarshal(data, &config)
}

// saveConfig saves the configuration to config.json.
func saveConfig() error {
    configMutex.RLock()
    data, err := json.MarshalIndent(config, "", "  ")
    configMutex.RUnlock()

    if err != nil {
        return err
    }

    return ioutil.WriteFile("config.json", data, 0644)
}
