package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func CreateDatabase(dbName string) error {
	createDB := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)
	_, err := DB.Exec(createDB)
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
		return err
	}

	return nil
}

// Connect 함수는 MySQL 데이터베이스에 연결합니다.
func Connect() error {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	// MySQL 데이터베이스 연결 문자열
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)

	// 데이터베이스 연결 시도
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
		return err
	}

	// 연결을 확인하기 위한 Ping
	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
		return err
	}

	log.Println("Database connected successfully")

	if err = CreateDatabase(dbname); err != nil {
		log.Fatalf("Failed to create database: %v", err)
		return err
	}

	// Users // 
	if err = CreateUserTable(); err != nil {
		log.Fatalf("Failed to create user table: %v", err)
		return err
	}

	// Courses //
	if err = CreateCoursesTable(); err != nil {
		log.Fatalf("Failed to create courses table: %v", err)
		return err
	}
	if err = CreateEnrollmentsTable(); err != nil {
		log.Fatalf("Failed to create enrollments table: %v", err)
		return err
	}

	return nil
}
