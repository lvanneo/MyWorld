package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"  //引入mysql驱动
)

func main() {
	/*DSN数据源名称
		[username[:password]@][protocol[(address)]]/dbname[?param1=value1&paramN=valueN]
		user@unix(/path/to/socket)/dbname 
		user:password@tcp(localhost:5555)/dbname?charset=utf8&autocommit=true 
		user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname?charset=utf8mb4,utf8 
		user:password@/dbname 
		无数据库: user:password@/
	*/
	db, err := sql.Open("mysql", "root:root@tcp(192.168.1.203:3306)/mytest") //第一个参数数驱动名
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 欲编译语句，插入数据，这个是标准的go接口，所以只要标准sql，其他数据库通用的，只要换上面的驱动名
	stmtIns, err := db.Prepare("INSERT INTO squareNum VALUES( ?, ? )" ) // ? = 占位符
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close() // main结束是关闭

	// Prepare statement for reading data
	stmtOut, err := db.Prepare("SELECT squareNumber FROM squareNum WHERE number = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	// Insert square numbers for 0-24 in the database
	for i := 0; i < 25; i++ {
		_, err = stmtIns.Exec(i, (i * i)) // 执行插入
		if err != nil {
			panic(err.Error())
		}
	}

	var squareNum int // we "scan" the result in here

	// Query the square-number of 13
	err = stmtOut.QueryRow(13).Scan(&squareNum) // WHERE number = 13
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("The square number of 13 is: %d", squareNum)

	// Query another number.. 1 maybe?
	err = stmtOut.QueryRow(1).Scan(&squareNum) // WHERE number = 1
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("The square number of 1 is: %d", squareNum)
}