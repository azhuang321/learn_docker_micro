package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"user/domain/repository"
	service2 "user/domain/service"
	"user/handler"
	user "user/proto/user"
)

func main() {
	//服务参数设置
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)
	//初始化服务
	srv.Init()

	//创建数据库连接
	db, err := gorm.Open("mysql", "root:123123@/micro?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.SingularTable(true)

	//只执行一次
	//rp := repository.NewUserRepository(db)
	//rp.InitTable()

	//创建服务实例
	userDataService := service2.NewUserDataService(repository.NewUserRepository(db))

	err = user.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService})

	if err != nil {
		fmt.Println(err)
	}

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
