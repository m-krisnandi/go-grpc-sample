package gapi

import (
	"context"
	"go-grpc-sample/model"
	pb "go-grpc-sample/proto/book"
	"go-grpc-sample/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookGapi struct {
	pb.UnimplementedBookGapiServer
	repository repository.BookRepository
}

func NewGrpcBook(repo repository.BookRepository) *BookGapi {
	return &BookGapi{repository: repo}
}

func (s *BookGapi) InsertBook(ctx context.Context, input *pb.InsertBookRequest) (*pb.InsertBookResponse, error) {
	data := model.Book{
		Title:       input.Title,
		Author:      input.Author,
		Description: input.Description,
		CreatedBy:   "system",
		UpdatedBy:   "system",
	}

	book, err := s.repository.InsertBook(data)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := pb.InsertBookResponse{
		Meta: &pb.Meta{
			Code:    201,
			Message: "Book: Insert Book",
		},
		Data: &pb.BookDBResponse{
			Id:          book.Id,
			Title:       book.Title,
			Author:      book.Author,
			Description: book.Description,
			CreatedAt:   book.CreatedAt,
			CreatedBy:   book.CreatedBy,
			UpdatedAt:   book.UpdatedAt,
			UpdatedBy:   book.UpdatedBy,
		},
	}

	return &resp, nil
}

func (s *BookGapi) GetListBooks(ctx context.Context, input *pb.GetListBooksRequest) (*pb.GetListBooksResponse, error) {
	filter := map[string]interface{} {
		"title": input.Title,
		"author": input.Author,
	}

	books, err := s.repository.GetListBooks(filter)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := []*pb.BookDBResponse{}
	for _, val := range books {
		book := pb.BookDBResponse{
			Id:          val.Id,
			Title:       val.Title,
			Author:      val.Author,
			Description: val.Description,
			CreatedAt:   val.CreatedAt,
			CreatedBy:   val.CreatedBy,
			UpdatedAt:   val.UpdatedAt,
			UpdatedBy:   val.UpdatedBy,
		}
		
		resp = append(resp, &book)
	}

	result := pb.GetListBooksResponse{
		Meta: &pb.Meta{
			Code: 200,
			Message: "Book: Get List Books",
		},
		Data: resp,
	}

	return &result, nil
}

func (s *BookGapi) GetBookById(ctx context.Context, input *pb.GetBookByIdRequest) (*pb.GetBookByIdResponse, error) {
    if input.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "book id is empty")
	}

	book, err := s.repository.GetBookById(input.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result := pb.GetBookByIdResponse{
		Meta: &pb.Meta{
			Code: 200,
			Message: "Book: Get Book By Id",
		},
		Data: &pb.BookDBResponse{
			Id:          book.Id,
			Title:       book.Title,
			Author:      book.Author,
			Description: book.Description,
			CreatedAt:   book.CreatedAt,
			CreatedBy:   book.CreatedBy,
			UpdatedAt:   book.UpdatedAt,
			UpdatedBy:   book.UpdatedBy,
		},
	}

	return &result, nil
}

func (s *BookGapi) UpdateBookById(ctx context.Context, input *pb.UpdateBookByIdRequest) (*pb.UpdateBookByIdResponse, error) {
	if input.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "book id is empty")
	}

	updateData := map[string]interface{}{
		"title":       input.Title,
        "author":      input.Author,
        "description": input.Description,
	}

	book, err := s.repository.UpdateBookById(input.Id, updateData)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result := pb.UpdateBookByIdResponse{
		Meta: &pb.Meta{
			Code: 200,
			Message: "Book: Update Book By Id",
		},
		Data: &pb.BookDBResponse{
			Id:          book.Id,
			Title:       book.Title,
			Author:      book.Author,
			Description: book.Description,
			CreatedAt:   book.CreatedAt,
			CreatedBy:   book.CreatedBy,
			UpdatedAt:   book.UpdatedAt,
			UpdatedBy:   book.UpdatedBy,
		},
	}

	return &result, nil
}

func (s *BookGapi) DeleteBookById(ctx context.Context, input *pb.DeleteBookByIdRequest) (*pb.DeleteBookByIdResponse, error) {
	if input.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "book id is empty")
	}

	err := s.repository.DeleteBookById(input.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result := pb.DeleteBookByIdResponse{
		Meta: &pb.Meta{
			Code: 200,
			Message: "Book: Delete By Id",
		},
	}

	return &result, nil
}