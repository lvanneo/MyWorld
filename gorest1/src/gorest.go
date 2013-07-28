package main

import (
    "fmt"
    "github.com/drone/routes"
    "net/http"
)

func getuser(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Get\n")
    params := r.URL.Query()
    uid := params.Get(":uid")
    fmt.Fprintf(w, "you are get user %s", uid)
}

func modifyuser(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Post\n")
    params := r.URL.Query()
    uid := params.Get(":uid")
    productName := params.Get("productName")
    fmt.Fprintf(w, "you are modify user %s -- %s", uid , productName)
    fmt.Printf("Post : %s  %s\n", uid, productName)
//    fmt.Printf(params)
}

func deleteuser(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Delete\n")
    params := r.URL.Query()
    uid := params.Get(":uid")
    fmt.Fprintf(w, "you are delete user %s", uid)
}

func adduser(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Put\n")
    params := r.URL.Query()
    uid := params.Get(":uid")
    fmt.Fprint(w, "you are add user %s", uid)
}

func query(w http.ResponseWriter, r *http.Request){
    fmt.Printf("Query\n")
	params := r.URL.Query()
    uid := params.Get(":uid")
    fmt.Fprintf(w, "query %s", uid)
}


func main() {
    mux := routes.New()
    mux.Get("/query/:uid", query)
    mux.Get("/user/:uid", getuser)
    mux.Post("/user/:uid", modifyuser)
    mux.Del("/user/:uid", deleteuser)
    mux.Put("/user/", adduser)
    http.Handle("/", mux)
    http.ListenAndServe(":8088", nil)
}
