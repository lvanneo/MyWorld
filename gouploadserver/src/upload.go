package main

import (
    "fmt"
    "io"
    "net/http"
    "log"
    "os"
    "time"
)

// 获取大小的接口
type Sizer interface {
    Size() int64
}

//处理文件上传的 Web服务方法 
func UploadServer(w http.ResponseWriter, r *http.Request) {
    if "POST" == r.Method {
    
        file, h, err := r.FormFile("userfile")
        if err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        
		//name := r.Form["uploadfile"]
		//获取文件名
        filename := h.Filename
        
        defer file.Close()
        f,err:=os.Create(filename)
        defer f.Close()
        io.Copy(f,file)

		fmt.Fprintf(w, "上传文件的大小为: %d", file.(Sizer).Size())
		fmt.Printf("%s  Size: %d KB  Name：%s\n", time.Now().Format("2006-01-02 15:04:05"), file.(Sizer).Size()/1024, filename)
//		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		
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
    http.HandleFunc("/upload", UploadServer)
    err := http.ListenAndServe(":8086", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
