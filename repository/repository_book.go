package repository

import (
	"errors"
	"go-grpc-sample/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookRepositoryImpl struct {
	db *gorm.DB
}

type BookRepository interface {
	InsertBook(book model.Book) (*model.Book, error)
	GetListBooks(filter map[string]interface{}) ([]model.Book, error)
	GetBookById(id string) (*model.Book, error)
	UpdateBookById(id string, input map[string]interface{}) (*model.Book, error)
	DeleteBookById(id string) error
}

func NewDataBook(db *gorm.DB) BookRepository {
	return &BookRepositoryImpl{db}
}

func (r *BookRepositoryImpl) InsertBook(book model.Book) (*model.Book, error) {
	err := r.db.Create(&book).Error
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *BookRepositoryImpl) GetListBooks(filter map[string]interface{}) ([]model.Book, error) {
	var books []model.Book

	query := r.db

	// Filter by title
	if filter["title"] != "" {
		query = query.Where("title LIKE ?", "%"+filter["title"].(string)+"%")
	}

	// Filter by author
	if filter["author"] != "" {
		query = query.Where("author LIKE ?", "%"+filter["author"].(string)+"%")
	}

	// Filter by category
	// if filter["category"] != "" {
	// 	query = query.Where("category_id = ?", filter["category"].(string))
	// }

	err := query.Find(&books).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return books, nil
}

func (r *BookRepositoryImpl) GetBookById(id string) (*model.Book, error) {
	var book model.Book
	err := r.db.Model(&model.Book{}).
		Where("id = ?", id).
		First(&book).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &book, nil
}

func (r *BookRepositoryImpl) UpdateBookById(id string, input map[string]interface{}) (*model.Book, error) {
	var book model.Book

	err := r.db.Model(&book).
		Where("id = ?", id).
		Clauses(clause.Returning{}).
		Updates(input). 
		Error

	if err != nil {
		return nil, err
	}

	r.db.First(&book, id)

	return &book, nil
}


func (r *BookRepositoryImpl) DeleteBookById(id string) error {
	err := r.db.
		Where("id = ?", id).
		Delete(&model.Book{}).
		Error

	if err != nil {
		return err
	}

	return nil
}
