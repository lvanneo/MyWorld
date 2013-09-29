//多协程文件传输服务端
//作者：李林
//邮箱：lvan_software@foxmail.com
//版本：1.0
//日期：2013-09-26
//对每个请求由一个单独的协程进行处理，文件接收完成后由一个协负责将所有接收的数据合并成一个有效文件

package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		// host   = "192.168.1.5"	//如果写locahost或127.0.0.1则只能本地访问。
		port = "9090"
		// remote = host + ":" + port

		remote = ":" + port //此方式本地与非本地都可访问
	)

	fmt.Println("服务器初始化... (Ctrl-C 停止)")

	lis, err := net.Listen("tcp", remote)
	defer lis.Close()

	if err != nil {
		fmt.Println("监听端口发生错误: ", remote)
		os.Exit(-1)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("客户端连接错误: ", err.Error())
			// os.Exit(0)
			continue
		}

		//调用文件接收方法
		go receiveFile(conn)
	}
}

/*
*	文件接收方法
*	2013-09-26
*	李林
*
*	con 连接成功的客户端连接
 */
func receiveFile(con net.Conn) {
	var (
		res          string
		tempFileName string                    //保存临时文件名称
		data         = make([]byte, 1024*1024) //用于保存接收的数据的切片
		by           []byte
		databuf      = bytes.NewBuffer(by) //数据缓冲变量
		fileNum      int                   //当前协程接收的数据在原文件中的位置
	)
	defer con.Close()

	fmt.Println("新建立连接: ", con.RemoteAddr())
	j := 0 //标记接收数据的次数
	for {
		length, err := con.Read(data)
		if err != nil {

			// writeend(tempFileName, databuf.Bytes())
			da := databuf.Bytes()
			// fmt.Println("over", fileNum, len(da))
			fmt.Printf("客户端 %v 已断开. %2d %d \n", con.RemoteAddr(), fileNum, len(da))
			return
		}

		if 0 == j {

			res = string(data[0:8])
			if "fileover" == res { //判断是否为发送结束指令，且结束指令会在第一次接收的数据中
				xienum := int(data[8])
				mergeFileName := string(data[9:length])
				go mainMergeFile(xienum, mergeFileName) //合并临时文件，生成有效文件
				res = "文件接收完成: " + mergeFileName
				con.Write([]byte(res))
				fmt.Println(mergeFileName, "文件接收完成")
				return

			} else { //创建临时文件
				fileNum = int(data[0])
				tempFileName = string(data[1:length]) + strconv.Itoa(fileNum)
				fmt.Println("创建临时文件：", tempFileName)
				fout, err := os.Create(tempFileName)
				if err != nil {
					fmt.Println("创建临时文件错误", tempFileName)
					return
				}
				fout.Close()
			}
		} else {
			// databuf.Write(data[0:length])
			writeTempFileEnd(tempFileName, data[0:length])
		}

		res = strconv.Itoa(fileNum) + " 接收完成"
		con.Write([]byte(res))
		j++
	}

}

/*
*	把数据写入指定的临时文件中
*	2013-09-26
*	李林
*
*	fileName	临时文件名
*	data 		接收的数据
 */
func writeTempFileEnd(fileName string, data []byte) {
	// fmt.Println("追加：", name)
	tempFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		// panic(err)
		fmt.Println("打开临时文件错误", err)
		return
	}
	defer tempFile.Close()
	tempFile.Write(data)
}

/*
*	根据临时文件数量及有效文件名称生成文件合并规则进行文件合并
*	2013-09-26
*	李林
*
*	connumber	临时文件数量
*	filename 	有效文件名称
 */
func mainMergeFile(connumber int, filename string) {

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("创建有效文件错误", err)
		return
	}
	defer file.Close()

	//依次对临时文件进行合并
	for i := 0; i < connumber; i++ {
		mergeFile(filename+strconv.Itoa(i), file)
	}

	//删除生成的临时文件
	for i := 0; i < connumber; i++ {
		os.Remove(filename + strconv.Itoa(i))
	}

}

/*
*	将指定临时文件合并到有效文件中
*	2013-09-26
*	李林
*
*	rfilename	临时文件名称
*	wfile	 	有效文件
 */
func mergeFile(rfilename string, wfile *os.File) {

	// fmt.Println(rfilename, wfilename)
	rfile, err := os.OpenFile(rfilename, os.O_RDWR, 0666)
	defer rfile.Close()
	if err != nil {
		fmt.Println("合并时打开临时文件错误:", rfilename)
		return
	}

	stat, err := rfile.Stat()
	if err != nil {
		panic(err)
	}

	num := stat.Size()

	buf := make([]byte, 1024*1024)
	for i := 0; int64(i) < num; {
		length, err := rfile.Read(buf)
		if err != nil {
			fmt.Println("读取文件错误")
		}
		i += length

		wfile.Write(buf[:length])
	}

}
