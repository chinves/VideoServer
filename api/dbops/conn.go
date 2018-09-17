package dbops

import (
	"database/sql"
	"fmt"
)

/**
初始化db
 */
var (
	dbConn *sql.DB
	err error
)
func init(){
	fmt.Println("get dbconn... init...")
	dbConn,err := sql.Open("mysql", "root:123456@/video_server?charset=utf8")
	if err !=nil{
		fmt.Println("get dbconn err ...")
	}
	fmt.Printf("db %v\n",dbConn)
}
