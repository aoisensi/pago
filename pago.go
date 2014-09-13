package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "strconv"
)

type Page struct {
    Path string
    Body []byte
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL.Path)
    path := r.URL.Path

    if path == "/favicon.ico" || path == "/robots.txt" {
       data, err :=loadSolidData(path)
       if err != nil {
           writeError(w, 404)
           return
       }
       w.Write(data)
    }

    if strings.ContainsRune(path, '.') {
        writeError(w, 403)
        return
    }

    //load page
    page, err := loadTextPage(path)
    if err != nil {
        writeError(w, 404)
        return
    }
    fmt.Fprintf(w, "%s", page.Body)
}

func writeError(w http.ResponseWriter, code int) {
    w.WriteHeader(code)
    fmt.Fprintf(w, "%s", strconv.Itoa(code))
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8000", nil)
}

func loadTextPage(path string) (*Page, error) {
    filename := "./page" + path + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{ Path: path, Body: body }, nil
}

func loadSolidData(file string) ([]byte, error) {
    filename := "." + file
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return data, nil
}
