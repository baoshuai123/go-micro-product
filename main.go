package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"taobao/product/common"
	"taobao/product/domain/repository"
	service2 "taobao/product/domain/service"
	"taobao/product/handler"

	product "taobao/product/proto/product"
)

func main() {
	//配置中心
	consulConfig,err := common.GetConsulConfig("127.0.0.1",8500,"/micro/config")
	if err != nil {
		log.Error(err)
	}
	//注册中心
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	//链路追踪
	t,io,err := common.NewTracer("go.micro.service.product",
		"localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	//数据库设置
	mysqlInfo := common.GetMysqlFromConsul(consulConfig,"mysql")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlInfo.User,
		mysqlInfo.Pwd,
		mysqlInfo.Host,
		mysqlInfo.Port,
		mysqlInfo.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		log.Error(err)
	}
	sqlDB,_ := db.DB()
	defer sqlDB.Close()
	rp := repository.NewProductRepository(db)
	_ = rp.InitTable()
	productDataService := service2.NewProductDataService(rp)
	// 1 创建服务
	service := micro.NewService(
		micro.Name("go.micro.service.product"),
		micro.Version("latest"),
		//设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8083"),
		//添加consul作为注册中心
		micro.Registry(consul),
		//绑定链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	//2初始化服务
	service.Init()

	// Register Handler
	err = product.RegisterProductHandler(service.Server(), &handler.Product{ProductDataService: productDataService})
	if err != nil {
		log.Error(err)
		return
	}
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
