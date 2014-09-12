package main

import (
    "net/http"
)

const (
    listURL = "/_list/"

    pagePath = "./page"
)

type Page struct {
    Path string
    Body []byte
}

func main() {
    http.HandleFunc(listURL, handlerList)
    //http.HandleFunc("/_tree", handlerSolid)
    http.HandleFunc("/favicon.ico", handlerSolid)
    http.HandleFunc("/robots.txt", handlerSolid)
    http.HandleFunc("/", handlerPage)
    http.ListenAndServe(":8000", nil)
}
