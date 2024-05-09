package service

import (
	"context"

	"github.com/yaninyzwitty/kafka-producer-go/model"
	"github.com/yaninyzwitty/kafka-producer-go/repository"
)

type ProductService interface {
	CreateProduct(ctx context.Context, product model.Product) (model.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) CreateProduct(ctx context.Context, product model.Product) (model.Product, error) {
	return s.repo.CreateProduct(ctx, product)

}
