package main

//Go 连接MySQL示例
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	dsn := "root:wj770826@tcp(127.0.0.1:3306)/curd"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		fmt.Print(err)
	}
	return
}
func query() {

}

type user struct {
	name string
	age  int
	id   int
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("connect ok!")
	}
	var u1 user
	querySql := `select id,name,age from user where id=1 ;`
	rowObj := db.QueryRow(querySql)
	rowObj.Scan(&u1.id, &u1.name, &u1.age)
	fmt.Println(u1) //test

}
