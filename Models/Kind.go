package Models

import (
	"goTest/Database"
	_ "goTest/Database"
	"fmt"
	"log"
	"strings"
)

//定义结构
type Kind struct {
	Id         int    `xorm:"pk autoincr"`
	KindName   string `form:"kindName" json:"kindName" xorm:"KindName"`
	IsEnable   int    `form:"isEnable" json:"isEnable" xorm:"IsEnable"`
	CreateTime string `xorm:"CreateTime"`
}

type KindService struct {
	Database.DbService
}
///获取总记录数
func (ts *KindService)KindCount() (counts int64) {
	var kind Kind
	ts.DbService.GetXorm()
	count, _ := ts.DbService.GetXorm().Count(&kind)
	counts = count
	return
}

///根据Id获取数据
func GetInfo(id int) []map[string]string {

	//var k [] Kind
	//DataBase.Db.SQL("SELECT Id, KindName, IsEnable, CreateTime  FROM Kind ").Find(&k)

	//err:=DataBase.Db.SQL("SELECT Id, KindName, IsEnable, CreateTime  FROM Kind ").Where(" Id in (1,2,3)").Find(&k)

	//var p= "1,2,3,4"
	//err := DataBase.Db.SQL("SELECT Id, KindName, IsEnable, CreateTime  FROM Kind where  Id in (" + p + ")").Find(&k)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}

	//whereStr := ""
	//ids := "1"
	//isEnable := "0"
	//var ps []string
	//ps = append(ps, "1")
	//ps = append(ps, "0")
	//if len(ids) > 0 {
	//	whereStr += " And id=?"
	//}
	//if len(isEnable) > 0 {
	//	whereStr += " ANd isEnable=? "
	//}

	//fmt.Println(whereStr)

	var k [] Kind
	//k.IsEnable = 3

	var para [] int
	para = append(para, 4)
	para = append(para, 5)
	para = append(para, 6)
	storeids:="1,2,3,4"


	Database.Db.SQL(" SELECT Id, KindName, IsEnable, CreateTime  FROM Kind where  Id = ? ",1).Find(&k)

	if len(k)>0 {

		fmt.Println("ok")
	}
	var math=strings.Split(storeids,",")

	var kmodel Kind
	kmodel.IsEnable=9
	r, err := Database.Db.In("Id", math).Update(&kmodel)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}

	fmt.Println(r)

	eng := Database.Db
	session := eng.NewSession()
	defer session.Close()
	session.Begin()
	user1 := Kind{KindName: "xiaoxiao1", IsEnable: 1, CreateTime: "2019-07-03 20:16:37"}
	_, err = session.Insert(&user1)
	if err != nil {
		session.Rollback()
	}
	err = session.Commit()
	if err != nil {
	}

	return nil
}

///获取所有数据
func GetAllKindData() (kinds [] Kind) {
	sql := " select Id, KindName, IsEnable, CreateTime from Kind where 1=1 "
	Database.Db.SQL(sql).Find(&kinds)
	return kinds
}

///新增数据
func AddKindData(kind Kind) (bool bool) {
	r, _ := Database.Db.Insert(&kind)
	if r > 0 {
		return true
	} else {
		return false
	}
}

///修改
func UpdateKindData(kind Kind) (bool bool) {
	r, err := Database.Db.Id(kind.Id).Update(&kind)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if r > 0 {
		return true
	} else {
		return false
	}
}

///获取统计条数
func (ts *KindService)GetQueryData() int {
	sql:="SELECT COUNT(1) AS totle  FROM kind WHERE 1=1 "
	stmt,err:=ts.GetConnNative().Prepare(sql)
	var params []interface{}
	sum,err:=stmt.Query(params...)

	if err!=nil {
		panic(err)
	}
	totle:=0
	sum.Next()
	sum.Scan(&totle)
	fmt.Println(sum)
	return totle
}

