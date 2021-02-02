package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"taobao/product/common"
	go_micro_service_product "taobao/product/proto/product"
)

func main()  {
	//注册中心
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	//链路追踪
	t,io,err := common.NewTracer("go.micro.service.product.client",
		"localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	service := micro.NewService(
		micro.Name("go.micro.service.product.client"),
		micro.Version("latest"),
		//设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8083"),
		//添加consul作为注册中心
		micro.Registry(consul),
		//绑定链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
	)
	productService:=go_micro_service_product.NewProductService("go.micro.service.product",service.Client())
	productAdd := &go_micro_service_product.ProductInfo{
		ProductName:        "github",
		ProductSku:         "马局长",
		ProductPrice:       998.8,
		ProductDescription: "github-马局长",
		ProductImage:       []*go_micro_service_product.ProductImage{
			{
				ImageName: "马局长-image",
				ImageCode: "马局长-image01",
				ImageUrl:  "马局长-image01",
			},
			{
				ImageName: "tom-image",
				ImageCode: "tom-image01",
				ImageUrl:  "tom-image01",
			},
		},
		ProductSize:        []*go_micro_service_product.ProductSize{
			{
				SizeName: "马局长-size",
				SizeCode: "马局长-size-code",
			},
		},
		ProductSeo:         &go_micro_service_product.ProductSeo{
			SeoTitle: "马局长-seo",
			SeoKeywords: "马局长-seo",
			SeoDescription: "马局长-seo",
			SeoCode: "马局长-seo",
		},
		ProductCategoryId:  1,
	}
	response, err := productService.AddProduct(context.TODO(), productAdd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
	
}
