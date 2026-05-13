package mapper

import (
	"gatewayapp/internal/dto"
	"myapp/api/gen/product"
)

func ToProductInfo(product *product.ProductInfo) dto.ProductInfo {
	if product == nil {
		return dto.ProductInfo{}
	}
	return dto.ProductInfo{
		ID:          product.Id,
		Name:        product.Name,
		Price:       product.Price,
		Stock:       product.Stock,
		StoreID:     product.StoreId,
		SKU:         product.Sku,
		Description: product.Description,
		Category:    product.Category,
		IsActive:    product.IsActive,
		ImageURL:    product.ImageUrl,
	}

}

func ToCreateProductRequest(req *dto.CreateProductRequest) *product.CreateProductRequest {
	return &product.CreateProductRequest{
		StoreId:     req.StoreID,
		Sku:         req.SKU,
		Name:        req.Name,
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
		IsActive:    req.IsActive,
		ImageUrl:    req.ImageURL,
		Description: req.Description,
	}
}

func ToCreateProductResponse(resp *product.CreateProductResponse) *dto.CreateProductResponse {
	return &dto.CreateProductResponse{
		Product: ToProductInfo(resp.Product),
	}
}

func ToUpdateProductRequest(req *dto.UpdateProductRequest) *product.UpdateProductRequest {
	return &product.UpdateProductRequest{
		Id:          req.ID,
		StoreId:     req.StoreID,
		Sku:         req.SKU,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
		IsActive:    req.IsActive,
		ImageUrl:    req.ImageURL,
	}
}

func ToUpdateProductResponse(resp *product.UpdateProductResponse) *dto.UpdateProductResponse {
	return &dto.UpdateProductResponse{
		Product: ToProductInfo(resp.Product),
	}
}

func ToDeleteProductRequest(req *dto.DeleteProductRequest) *product.DeleteProductRequest {
	return &product.DeleteProductRequest{
		StoreId:   req.StoreID,
		UserId:    req.UserID,
		ProductId: req.ProductID,
	}
}
func ToDeleteProductResponse(resp *product.DeleteProductResponse) *dto.DeleteProductResponse {
	return &dto.DeleteProductResponse{
		Success: resp.Success,
	}
}

func ToGetProductByIDRequest(req *dto.GetProductByIDRequest) *product.GetProductByIDRequest {
	return &product.GetProductByIDRequest{
		StoreId:   req.StoreID,
		ProductId: req.ProductID,
	}
}

func ToGetProductByIDResponse(req *product.GetProductResponse) *dto.GetProductResponse {
	return &dto.GetProductResponse{
		Product: ToProductInfo(req.Product),
	}
}
