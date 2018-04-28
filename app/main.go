package app

import (
    "net/http"

    "github.com/gorilla/mux"
    "github.com/inct07/kawaii-api/gae"
)

func init() {
    r := mux.NewRouter()
    r.HandleFunc("/", gae.SayHello)
    http.Handle("/", r)
}
