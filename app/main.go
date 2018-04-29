package app

import (
    "net/http"

    "github.com/gorilla/mux"
    "encoding/json"
    "fmt"
)

func init() {
    r := mux.NewRouter()
    r.HandleFunc("/users/{id:[0-9]+}", UserHandler).Methods("GET")
    http.Handle("/", r)
}

type Reply struct {
    Message string `json:"message"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    reply := Reply {
        Message: "id(" + id + ")を取得しました",
    }

    json, _ := json.Marshal(reply)

    fmt.Fprint(w, string(json))
}