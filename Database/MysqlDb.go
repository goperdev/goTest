package Database

import (
	"goTest/Common"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"fmt"
)

type DbService struct {

}

var Db *xorm.Engine
///初始化加载
func init() {
	micro := Common.Microservice{}
	micro.LoadConfig()

	var err error
	params:=fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=true",
		micro.Mysql.User,
		micro.Mysql.Pwd,
		micro.Mysql.Host,
		micro.Mysql.Port,
		micro.Mysql.Default)

	Db, err = xorm.NewEngine("mysql", params)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	err = Db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (b *DbService)GetXorm()*xorm.Engine {
	micro := Common.Microservice{}
	micro.LoadConfig()

	params := fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=true",
		micro.Mysql.User,
		micro.Mysql.Pwd,
		micro.Mysql.Host,
		micro.Mysql.Port,
		micro.Mysql.Default)
	var err error
	engine, err := xorm.NewEngine("mysql", params)
	engine.ShowSQL()

	err = engine.Ping()
	if err != nil {
		panic(err)
	}
	return engine
}

func (b *DbService)GetConnNative() *sql.DB {
	micro :=Common.Microservice{}
	micro.LoadConfig()
	params := fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=true",
		micro.Mysql.User,
		micro.Mysql.Pwd,
		micro.Mysql.Host,
		micro.Mysql.Port,
		micro.Mysql.Default)
	var err error
	engine, err := sql.Open("mysql", params)
	if err != nil {
		panic(err)
	}
	err = engine.Ping()
	return engine
}

///redis绑定
func (b *DbService) Redis()  {

}