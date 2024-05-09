package repository

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/yaninyzwitty/kafka-producer-go/model"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product model.Product) (model.Product, error)
}

type productRepository struct {
	ctx    context.Context
	writer *kafka.Writer
}

func NewProductRepository(ctx context.Context, writer *kafka.Writer) ProductRepository {
	return &productRepository{ctx, writer}
}

func (r *productRepository) CreateProduct(ctx context.Context, product model.Product) (model.Product, error) {
	product.ID = uuid.New()
	productInJson, _ := json.Marshal(product)
	err := r.writer.WriteMessages(ctx, kafka.Message{
		Value: productInJson,
	})

	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}
