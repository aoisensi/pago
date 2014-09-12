package main

import (
    "io/ioutil"
    "net/http"
)

type Page struct {
    Path string
    Body []byte
}

func main() {
    http.HandleFunc("/favicon.ico", handlerSolid)
    http.HandleFunc("/robots.txt", handlerSolid)
    http.HandleFunc("/", handlerPage)
    http.ListenAndServe(":8000", nil)
}

func loadSolidData(file string) ([]byte, error) {
    filename := "." + file
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return data, nil
}
