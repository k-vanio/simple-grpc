package service

import (
	"context"

	"github.com/k-vanio/simple-grpc/internal/db/models"
	"github.com/k-vanio/simple-grpc/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryModel models.Category
}

func NewCategoryService(categoryModel models.Category) *CategoryService {
	return &CategoryService{
		CategoryModel: categoryModel,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryModel.Create(in.Name, in.Description)
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, err.Error())
	}

	response := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{Category: response}, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Black) (*pb.CategoryList, error) {
	categories, err := c.CategoryModel.FindAll()
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, err.Error())
	}

	var categoryList []*pb.Category
	for _, row := range categories {
		category := &pb.Category{
			Id:          row.ID,
			Name:        row.Name,
			Description: row.Description,
		}

		categoryList = append(categoryList, category)
	}

	return &pb.CategoryList{Categories: categoryList}, nil
}
