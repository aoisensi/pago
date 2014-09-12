package main

import (
    "net/http"
    "io/ioutil"
)

func handlerPage(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path

    if path[len(path) - 1] == '/' {
        path += "index"
    }
    page, err := loadPage(path)
    if err != nil {
        w.WriteHeader(404)
        w.Write([]byte("404"))
        return
    }
    w.Write(page.Body)
}

func loadPage(path string) (*Page, error) {
    path = pagePath + path
    page, err := loadTextPage(path)
    if page != nil {
        return page, nil
    }
    return nil, err
}

func loadTextPage(path string) (*Page, error) {
    filename := path + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{ Path: path, Body: body }, nil
}
