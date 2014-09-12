package main

import (
    "net/http"
    "io/ioutil"
)

func handlerSolid(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[1:]

    data, err := ioutil.ReadFile(path)
    if err != nil {
        w.WriteHeader(404)
        w.Write([]byte("404"))
        return
    }

    w.Write(data)
}
