package main
      
import (
    "fmt"
    "log"
    "net"
    "bufio"
)
      
func handleConnection(conn net.Conn) {
	for {
	    data, err := bufio.NewReader(conn).ReadString('\n')
	    if err != nil {
	        log.Fatal("get client data error: ", err)
	    }
	    
	    fmt.Printf("%#v\n", data)
	      
	    fmt.Fprintf(conn, "hello client\n")
	    
	    if data[0] == 48 {
	    	fmt.Fprintf(conn, "连接关闭！\n")
	    	conn.Close()
	    	break
	    }   
    }
}
      
func main() {
    ln, err := net.Listen("tcp", ":6010")
    if err != nil {
        panic(err)
    }
      
    for {
        conn, err := ln.Accept()
        if err != nil {
            log.Fatal("get client connection error: ", err)
        }
      
        go handleConnection(conn)
    }
}

//该代码片段来自于: http://www.sharejs.com/codes/go/4378