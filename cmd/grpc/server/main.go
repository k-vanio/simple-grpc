package main

import (
	"database/sql"
	"net"

	"github.com/k-vanio/simple-grpc/internal/db/models"
	"github.com/k-vanio/simple-grpc/internal/pb"
	"github.com/k-vanio/simple-grpc/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryModel := models.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryModel)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
