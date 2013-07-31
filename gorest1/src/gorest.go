package main

import (
    "fmt"
    "github.com/drone/routes"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type InfoObject struct{

	ProductName string
	
	Price float32
}

func getuser(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    uid := params.Get(":uid")
    fmt.Fprintf(w, "you are get user %s", uid)
    
    fmt.Printf("Get  %s \n", uid)
}

func modifyuser(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    uid := params.Get(":uid")
    productName := params.Get("ProductName")
    fmt.Fprintf(w, "you are modify user %s -- %s", uid , productName)
    
    defer r.Body.Close()
    input,err:=ioutil.ReadAll(r.Body)
    if err != nil {
    	fmt.Printf("error")
    }
    fmt.Printf("Post  %s  %s\n", uid, input)
//    fmt.Printf("input : %s \n",input)
//    fmt.Printf("%#v\n", input[0])
    var sss []byte = input[2:]
    //sss = "%7B%22ProductName%22%3A%22yuanzhifei89%22%2C%22Price%22%3A15.3%7D"
    
    
    var jsonInfo InfoObject
    json.Unmarshal(sss,&jsonInfo)
    fmt.Printf(jsonInfo.ProductName)
    fmt.Printf("%s", sss)
    
}

func deleteuser(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    uid := params.Get(":uid")
    fmt.Fprintf(w, "you are delete user %s", uid)
    fmt.Printf("Delete  %s \n", uid)
}

func adduser(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()
    uid := params.Get(":uid")
    fmt.Fprint(w, "you are add user %s", uid)
    
    input,err:=ioutil.ReadAll(r.Body)
    if err != nil {
    	fmt.Printf("error")
    }
    fmt.Printf("Put %s %s \n", uid, input)
//    fmt.Printf("input : %s \n",input)
    
//    fmt.Printf("%#v\n", input[0])
    
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
