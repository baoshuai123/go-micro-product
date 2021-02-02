package handler

import (
	"context"
	"taobao/product/common"
	"taobao/product/domain/model"
	"taobao/product/domain/service"
	. "taobao/product/proto/product"
)
type Product struct{
     ProductDataService service.IProductDataService
}
// AddProduct 添加商品
func (p *Product) AddProduct(ctx context.Context,request *ProductInfo,response *ResponseProduct) error  {
	productAdd := &model.Product{}
	if err := common.SwapTo(request,productAdd);err!=nil{
		return err
	}
	productID,err := p.ProductDataService.AddProduct(productAdd)
	if err != nil {
		return err
	}
	response.ProductId = productID
	return nil
}
// FindProductByID 根据id查找商品
func (p *Product) FindProductByID(ctx context.Context,request *RequestID,response *ProductInfo) error  {
	productData,err := p.ProductDataService.FindProductByID(request.ProductId)
	if err != nil {
		return err
	}
	if err := common.SwapTo(productData,response);err!=nil{
		return err
	}
	return nil
}
// UpdateProduct 商品更新
func (p *Product) UpdateProduct(ctx context.Context,request *ProductInfo,response *Response) error  {
	productUpdate := &model.Product{}
	if err := common.SwapTo(request,productUpdate);err !=nil{
		return err
	}
	err := p.ProductDataService.UpdateProduct(productUpdate)
	if err != nil {
		return err
	}
	response.Msg = "更新成功"
	return nil
}
// DeleteProductByID 根据id删除商品
func (p *Product) DeleteProductByID(ctx context.Context,request *RequestID,response *Response) error {
	if err := p.ProductDataService.DeleteProduct(request.ProductId);err != nil {
		return err
	}
	response.Msg = "删除成功"
	return nil
}

func (p *Product) FindAllProduct(ctx context.Context,request *RequestAll,response *AllProduct) error  {
	productAll,err := p.ProductDataService.FindAllProduct()
	if err != nil {
		return err
	}
	for _, v := range productAll {
		productInfo := &ProductInfo{}
		err := common.SwapTo(v,productInfo)
		if err != nil {
			return err
		}
		response.ProductInfo =append( response.ProductInfo,productInfo)
	}
	return nil
}
