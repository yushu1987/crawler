package main

import (
	"crawler/database/es"
	"crawler/database/mysql"
	"crawler/database/redis"
	"crawler/lib"
	"crawler/service"
)

func main() {
	lib.InitConfig()
	lib.InitLog()
	redis.InitRedis()
	es.InitEsConn()
	mysql.InitMysqlConn()
	/*
	var esclient = &es.ES{"crawler", "title"}
	es.InitEsConn()

	obj:=structs.Link{
		Id:1,
		Title:"台湾终将统一",
		Url:"http://www.baidu.com",
		Size:1000,
	}
	esclient.PutObject(obj,"1")

	l:=esclient.SearchTitle("台湾")
	for _ ,i :=range l {
		fmt.Println(i)
	}


	p:=&lib.ParseUrl{}
	p.NewDocument("http://www.sina.com.cn")
	t:=p.GetTitle()
	fmt.Println("title",t)
	p.GetImages()
	*/
/*
	service.Download()
	service.PutES()
	service.MainHandler()
*/

	service.StartServer()
}
