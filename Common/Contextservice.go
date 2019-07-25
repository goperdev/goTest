package Common

import 	(
	"github.com/BurntSushi/toml"
	"log"
)


type Microservice struct {
	Service struct {
		//服务英文名
		Name string
		//服务中文名
		DisplayName string
		//版本号
		Version string
		Host    string
		Port    int
	}

	//Consul服务地址
	Consul struct {
		Host string
		Port int
	}

	//数据库链接字符串
	Mysql struct {
		Host    string
		Port    int
		User    string
		Pwd     string
		Default string
	}

	//缓存链接字符串
	Cache struct {
		Host    string
		Port    int
		Pwd     string
		Default int
	}

	Rsa struct {
		PublicKey  string
		PrivateKey string
		Issuer     string
	}

	//Route *gin.Engine
}

func (m *Microservice) LoadConfig() (service Microservice) {
	_, err := toml.DecodeFile("appsetting.toml", &m)
	if err != nil {
		log.Fatal(err)
	}
	return *m
}