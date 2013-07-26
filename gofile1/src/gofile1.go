package main
import (
        "os"
        "fmt"
        "io"
        "bufio"
        "io/ioutil"
        "path/filepath"
        "flag"
)


//写文件
func main2() {
        userFile := "test.txt"
        fout,err := os.Create(userFile)
        defer fout.Close()
        if err != nil {
                fmt.Println(userFile,err)
                return
        }
        for i:= 0;i<10;i++ {
                fout.WriteString("Just a string test!\r\n")
                fout.Write([]byte("Just a []byte test!\r\n"))
        }
}

//读文件
func main3() {
        userFile := "test.txt"
        fin,err := os.Open(userFile)
        defer fin.Close()
        if err != nil {
                fmt.Println(userFile,err)
                return
        }
        buf := make([]byte, 1024)
        for{
                n, _ := fin.Read(buf)
                if 0 == n { 
                	break 
                }
                os.Stdout.Write(buf[:n])
        }
}


//删除文件.使用os库os.Open os.Create
func main4() {
    fi, err := os.Open("input.txt")
    if err != nil { panic(err) }
    defer fi.Close()
 
    fo, err := os.Create("output.txt")
    if err != nil { panic(err) }
    defer fo.Close()
 
    buf := make([]byte, 1024)
    for {
        n, err := fi.Read(buf)
        if err != nil && err != io.EOF { panic(err) }
        if n == 0 { break }
 
        if n2, err := fo.Write(buf[:n]); err != nil {
            panic(err)
        } else if n2 != n {
            panic("error in writing")
        }
    }
}

//删除文件.使用bufio库
func main5() {
    fi, err := os.Open("input.txt")
    if err != nil { panic(err) }
    defer fi.Close()
    r := bufio.NewReader(fi)
 
    fo, err := os.Create("output.txt")
    if err != nil { panic(err) }
    defer fo.Close()
    w := bufio.NewWriter(fo)
 
    buf := make([]byte, 1024)
    for {
        n, err := r.Read(buf)
        if err != nil && err != io.EOF { panic(err) }
        if n == 0 { break }
 
        if n2, err := w.Write(buf[:n]); err != nil {
            panic(err)
        } else if n2 != n {
            panic("error in writing")
        }
    }
 
    if err = w.Flush(); err != nil { panic(err) }
}

//删除文件.使用ioutil库
func main6() {
    b, err := ioutil.ReadFile("input.txt")
    if err != nil { panic(err) }
 
    err = ioutil.WriteFile("output.txt", b, 0644)
    if err != nil { panic(err) }
}


func getFilelist(path string) {
        err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
                if ( f == nil ) {return err}
                if f.IsDir() {return nil}
                println(path)
                return nil
        })
        if err != nil {
                fmt.Printf("filepath.Walk() returned %v\n", err)
        }
}

//遍历文件夹
func main7(){
        flag.Parse()
        root := flag.Arg(0)
        getFilelist(root)
}

func main(){
	
	main2()
	
}
