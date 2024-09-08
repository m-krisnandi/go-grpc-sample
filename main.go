package main

import (
	"go-grpc-sample/initialization"
	"go-grpc-sample/model"
)

func main() {
	db, err := initialization.DbInit()
	if err != nil {
		panic(err.Error())
	}

	_ = db.AutoMigrate(&model.Book{})

	err = initialization.ServerInit(db)
	if err != nil {
		panic(err)
	}
}