package main

import (
    "fmt"
    "github.com/drone/routes"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "strings"
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
    
//    fmt.Printf("Post  %s  %s\n", uid, input)
//    fmt.Printf("input : %s \n",input)
//    fmt.Printf("%#v\n", input[0])
    var sss []byte = input[2:]
    
//    fmt.Println(sss) 
    
    jstr := URLJsonDecoder(string(sss))
    
    var jsonInfo InfoObject
    json.Unmarshal([]byte(jstr),&jsonInfo)
    fmt.Println(jstr)
    fmt.Printf("name:  %s \n", jsonInfo.ProductName)
    fmt.Printf("price:  %f \n", jsonInfo.Price)
    fmt.Println("price: ",jsonInfo.Price)
    
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

func URLJsonDecoder(jsonStr string) (json string){
	jsonStr = strings.Replace(jsonStr, "%7B", "{" , -1)
	jsonStr = strings.Replace(jsonStr, "%7D", "}" , -1)
	jsonStr = strings.Replace(jsonStr, "%22", "\"" , -1)
	jsonStr = strings.Replace(jsonStr, "%3A", ":" , -1)
	jsonStr = strings.Replace(jsonStr, "%2C", "," , -1)
	
	return jsonStr
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
