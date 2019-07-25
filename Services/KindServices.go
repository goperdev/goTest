package Services

import (
	"goTest/Common"
	. "goTest/Models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/guregu/null.v3"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Test struct {
	Id int
	Name string
	Date time.Time `xorm:"default 'NULL' comment('最后更新时间') TIMESTAMP" time_format:"2006-01-02 15:04:05"`
	//Date time.Time `xorm:"default 'NULL' comment('最后更新时间') TIMESTAMP" `
}

///接口首页
func Index(c* gin.Context) {
	//c.String(http.StatusOK, "Video接口数据中心")
	test:=Test{}
	test.Id=1
	test.Name="测试"
	const format = "2006-01-02 15:04:05"

	//test.Date,_=time.Parse(format,"2019-07-10 14:33:55")
	//t:=time.Now().Local()
	c.JSON(http.StatusOK,test)
}

//获取总条数
func GetKindCount(c*gin.Context)  {
	//count:=KindCount()

	c.JSON(http.StatusOK,gin.H{
		"总记录数：":1,
	})
}

//根据Id获取数据
func GetKindId(c*gin.Context)() {
	id := c.Param("id")
	kid, _ := strconv.Atoi(id)
	user := &Kind{}
	c.ShouldBindWith(user, binding.Form)
	kind := GetInfo(kid)
	fmt.Print("user=",user)
	c.Header("X-Pagination-PageCount", "10")
	c.JSON(http.StatusOK, gin.H{
		"kind": kind,
	})
}

///获取所有数据
func GetAllKind(c* gin.Context)() {

	redis := Common.NewRedis()

	//redis.Client.Do("APPEND","test","kkk")
	var redisList = redis.Client.LLen("kindList")
	//判断redis是否有值,如果有值就取redis，没有就查数据库并把查询结构写入redis
	if redisList.Val() > 0 {

		var kindList []Kind
		var model Kind

		for i:=0;i< int(redisList.Val()) ;i++ {
			var rediKind=redis.Client.LIndex("kindList",int64(i))//获取对应下标的值字符串
			json.Unmarshal([]byte(rediKind.Val()),&model)
			kindList= append(kindList, model)
		}
		c.JSON(http.StatusOK,kindList)
	} else {
		kk := GetAllKindData()
		for v:=range kk  {
			jsoninfo,_:=json.Marshal(kk[v])
			redis.Client.RPush("kindList",jsoninfo)
		}
		kindmap := make([]Kind, 0)
		index := c.Request.FormValue("index")
		size := c.Request.FormValue("size")
		for v := range kk {
			kindmap = append(kindmap, kk[v])
		}
		pageIndex, _ := strconv.Atoi(index)
		pageSize, _ := strconv.Atoi(size)
		pageSize = pageSize * pageIndex
		pageIndex = (pageIndex - 1) * pageSize
		d := kindmap[pageIndex:pageSize]
		c.JSON(http.StatusOK, gin.H{
			"d": d,
		})
	}
}

///新增数据
func AddKind(c*gin.Context)() {



	user:= &Kind{}
	c.ShouldBindWith(user, binding.Form)

	kindName := c.Request.FormValue("kindName")
	isEnable := c.Request.FormValue("isEnable")
	IsEnable, err := strconv.Atoi(isEnable)
	if err != nil {
		log.Fatalln(err.Error())
	}

	datetime:=time.Now().Format("2006-01-02 15:04:05")
	user.CreateTime=datetime
	k := Kind{KindName: kindName, IsEnable: IsEnable,CreateTime:datetime}
	fmt.Println(k)




	result := AddKindData(*user)
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

//修改数据
func UpdateKind(c*gin.Context) {
	kindName := c.Request.FormValue("kindName")
	isEnable := c.Request.FormValue("isEnable")
	kid := c.Request.FormValue("id")

	IsEnable, err := strconv.Atoi(isEnable)
	if err != nil {
		log.Fatalln(err.Error())
	}
	Id, err := strconv.Atoi(kid)
	if err != nil {
		log.Fatalln(err.Error())
	}
	datetime:=time.Now().Format("2006-01-02 15:04:05")
	k := Kind{Id: Id, KindName: kindName, IsEnable: IsEnable,CreateTime:datetime}
	result := UpdateKindData(k)
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

///查询数据  返回string
func QueryDataString(c*gin.Context) {

	var kind KindService
	result := kind.GetQueryData()
	jsonestr, err := json.Marshal(result)
	if err != nil {
		c.String(http.StatusInternalServerError, "JSON转换错误")
		return
	}
	c.String(http.StatusOK,string(jsonestr))
}


///查询数据  返回string
func QueryData(c*gin.Context) {
	var kind KindService
	result := kind.GetQueryData()
	c.JSON(http.StatusOK, result)
}

func Tk(c*gin.Context) {
	params := &Tkinfo{}
	redis := Common.NewRedis()
	println("redis服务：", redis)
	if err := c.ShouldBindWith(params, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, "参数不正确")
		return
	}
	var redisValuelen = redis.Client.LLen("testData")//获取对应的键的值有多少条
	fmt.Println("redisValue=", redisValuelen)
	var mlist []Tkinfo
	for i := 0; i < int(redisValuelen.Val()); i++ {
		index := i
		var hh= redis.Client.LIndex("testData", int64(index))


		mm:=Tkinfo{}
		err:=json.Unmarshal([]byte(hh.Val()),&mm)
		if err != nil {
			c.String(http.StatusInternalServerError, "JSON转换错误")
			return
		}

		mlist= append(mlist, mm)
	}
	c.JSON(http.StatusOK, mlist)
}


type Tkinfo struct {
	Id     int    `xorm:"default 'NULL' comment('Id') INT(11)"`
	Sname  string `xorm:"default 'NULL' comment('Sname') VARCHAR(80)"`
	Remark null.String `xorm:"default 'NULL' comment('Remark') VARCHAR(80)"`

}


