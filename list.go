package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "path/filepath"
)

func handlerList(w http.ResponseWriter, r *http.Request) {
    path := pagePath + r.URL.Path[len(listURL) - 1:]
    fmt.Println(path);
    files, err := ioutil.ReadDir(path)
    if err != nil {
        w.WriteHeader(404)
        w.Write([]byte("404"))
        return
    }

    fmt.Fprintf(w, "<html><body><ul>")
    for _, file := range files {
        name := file.Name()
        name = name[:len(name) - len(filepath.Ext(name))]
        link := r.URL.Path[len(listURL) - 1:] + name
        if file.IsDir() {
            fmt.Fprintf(w, "<li><a href=\"%s/\">%s/</a></li>", link, name)
        } else {
            fmt.Fprintf(w, "<li><a href=\"%s\">%s</a></li>", link, name)
        }
    }
    fmt.Fprintf(w, "</ul></body></html>")
}
