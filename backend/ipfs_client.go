package main

import (
    "fmt"
    "io/ioutil"
    "mime/multipart"
    "net/http"
    "os"
)

func uploadFileToIPFS(file *multipart.FileHeader, nodeURL string) (string, error) {
    f, err := file.Open()
    if err != nil {
        return "", err
    }
    defer f.Close()

    tempFile, err := ioutil.TempFile(os.TempDir(), "upload-*.bin")
    if err != nil {
        return "", err
    }
    defer os.Remove(tempFile.Name())

    fileBytes, err := ioutil.ReadAll(f)
    if err != nil {
        return "", err
    }

    if _, err := tempFile.Write(fileBytes); err != nil {
        return "", err
    }

    res, err := http.Post(fmt.Sprintf("%s/api/v0/add", nodeURL), "multipart/form-data", tempFile)
    if err != nil {
        return "", err
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return "", err
    }

    // Parse the response to extract the CID
    // Assume response contains a field "Hash" with the CID
    var cid string
    // Parsing logic here...

    return cid, nil
}

func fetchFromIPFS(cid string, nodeURL string) (string, error) {
    res, err := http.Get(fmt.Sprintf("%s/ipfs/%s", nodeURL, cid))
    if err != nil {
        return "", err
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        return "", err
    }

    return string(body), nil
}
