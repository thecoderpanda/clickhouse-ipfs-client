package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

// Example function to upload a file to IPFS.
func uploadToIPFS(fileData []byte) (string, error) {
    url := fmt.Sprintf("%s/api/v0/add", config.IPFSNodeURL)
    req, err := http.NewRequest("POST", url, nil)
    if err != nil {
        return "", err
    }
    req.SetBasicAuth(config.IPFSUsername, config.IPFSPassword)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(body), nil // Parse and return CID from response.
}

// Example function to fetch data from IPFS.
func fetchFromIPFS(cid string) ([]byte, error) {
    url := fmt.Sprintf("%s/ipfs/%s", config.IPFSNodeURL, cid)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    return ioutil.ReadAll(resp.Body)
}
