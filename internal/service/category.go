package service

import (
	"github.com/mdvedovato/14-gRPC/internal/database"
	"github.com/mdvedovato/14-gRPC/internal/pb"
)

type CategoryService struct{
	pb.mustEmbedUnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService (categoryDB database.Category) *CategoryService {
	return &CategoryService {
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != null {
		return nil, err
	}

	categoryResponse := $pb.Category {
		Id: 			category.ID,
		Name: 			category.Name,
		Description:	category.Description,
	}

	return &pb.categoryResponse {
		Category: categoryResponse,
	}, nil
}