package main
   
import (
    "fmt"
    "net"
    "bufio"
)
   
func main() {
    conn, err := net.Dial("tcp", ":6010")
    if err != nil {
        panic(err)
    }
   
    fmt.Fprintf(conn, "hello server\n")
   
    data, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        panic(err)
    }
   
    fmt.Printf("%#v\n", data)
}
//该代码片段来自于: http://www.sharejs.com/codes/go/4378