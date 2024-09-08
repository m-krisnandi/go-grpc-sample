package initialization

import (
	"fmt"
	"go-grpc-sample/gapi"
	"go-grpc-sample/proto/book"
	"go-grpc-sample/repository"
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func init() {
    if err := godotenv.Load(); err != nil {
        fmt.Println("Error loading .env file")
    }
}

func DbInit() (*gorm.DB, error) {
	var dialect gorm.Dialector

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_SSLMODE"),
    )

	dialect = postgres.Open(dsn)

	db, err := gorm.Open(dialect, gormConfig)
	if err != nil {
		return nil, fmt.Errorf("error to open: %v", err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error: %v", err.Error())
	}

	// Maximum number idle connection
	sqlDB.SetMaxIdleConns(20)

	// Maximum number open connection
	sqlDB.SetMaxOpenConns(100)

	tm := time.Minute * time.Duration(20)

	sqlDB.SetConnMaxLifetime(tm)

	return db, nil
}

func ServerInit(db *gorm.DB) error {
	network := "tcp"
	address := ":1200"
	listen, err := net.Listen(network, address)
	if err != nil {
		return fmt.Errorf("error to listen: %v", err.Error())
	}

	defer listen.Close()

	grpcServer := grpc.NewServer()

	// register
	bookSvc := repository.NewDataBook(db)
	bookServer := gapi.NewGrpcBook(bookSvc)

	// put register server
	book.RegisterBookGapiServer(grpcServer, bookServer)

	fmt.Printf("Server started at http://localhost:%s\n", address)

	err = grpcServer.Serve(listen)
	if err != nil {
		return fmt.Errorf("error to serve: %v", err.Error())
	}

	return nil
}