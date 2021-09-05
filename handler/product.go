package handler

import (
	context "context"
	"github.com/wangjiandev/product/common"
	"github.com/wangjiandev/product/domain/model"
	"github.com/wangjiandev/product/domain/service"
	product "github.com/wangjiandev/product/proto/product"
)

type Product struct {
	ProductDataService service.IProductDataService
}

// AddProduct 新增商品
func (p *Product) AddProduct(ctx context.Context, info *product.ProductInfo, responseProduct *product.ResponseProduct) error {
	productAdd := &model.Product{}
	if err := common.SwapTo(info, productAdd); err != nil {
		return err
	}
	productId, err := p.ProductDataService.AddProduct(productAdd)
	if err != nil {
		return err
	}
	responseProduct.ProductId = productId
	return nil
}

// FindProductById 根据id查询商品
func (p *Product) FindProductById(ctx context.Context, requestId *product.RequestId, info *product.ProductInfo) error {
	productData, err := p.ProductDataService.FindProductByID(requestId.ProductId)
	if err != nil {
		return err
	}
	err = common.SwapTo(productData, info)
	if err != nil {
		return err
	}
	return nil
}

// UpdateProduct 更新
func (p *Product) UpdateProduct(ctx context.Context, info *product.ProductInfo, response *product.Response) error {
	productData := &model.Product{}
	err := common.SwapTo(info, productData)
	if err != nil {
		return err
	}
	err = p.ProductDataService.UpdateProduct(productData)
	if err != nil {
		return err
	}
	response.Message = "更新成功"
	return nil
}

// DeleteProduct 根据id删除对应商品以及相关信息
func (p *Product) DeleteProduct(ctx context.Context, requestId *product.RequestId, response *product.Response) error {
	err := p.ProductDataService.DeleteProduct(requestId.ProductId)
	if err != nil {
		return err
	}
	response.Message = "删除成功"
	return nil
}

// FindAllProduct 查询所有商品信息
func (p *Product) FindAllProduct(ctx context.Context, all *product.RequestAll, response *product.ResponseProductList) error {
	allProduct, err := p.ProductDataService.FindAllProduct()
	if err != nil {
		return err
	}
	for _, v := range allProduct {
		productInfo := &product.ProductInfo{}
		err := common.SwapTo(v, productInfo)
		if err != nil {
			return err
		}
		response.ProductInfoList = append(response.ProductInfoList, productInfo)
	}
	return nil
}
