package main

import (
    "fmt"
    "io"
    "net/http"
    "log"
    "os"
)

// 获取大小的借口
type Sizer interface {
    Size() int64
}

// hello world, the web server
func HelloServer(w http.ResponseWriter, r *http.Request) {
    if "POST" == r.Method {
    
        file, h, err := r.FormFile("userfile")
        if err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        
//    	name := r.Form["uploadfile"]
        filename := h.Filename
        
        defer file.Close()
        f,err:=os.Create(filename)
        defer f.Close()
        io.Copy(f,file)
		fmt.Fprintf(w, "上传文件的大小为: %d", file.(Sizer).Size())
		fmt.Printf("上传文件的大小为: %d  名称：%s\n", file.(Sizer).Size(), filename)
		
        return
    }

    // 上传页面
    w.Header().Add("Content-Type", "text/html")
    w.WriteHeader(200)
    html := `
<form enctype="multipart/form-data" action="/hello" method="POST">
    Send this file: <input name="userfile" type="file" />
    <input type="submit" value="Send File" />
</form>
`
    io.WriteString(w, html)
}

func main() {
    http.HandleFunc("/hello", HelloServer)
    err := http.ListenAndServe(":8086", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
