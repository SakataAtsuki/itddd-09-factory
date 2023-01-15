package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/SakataAtsuki/itddd-09-factory/domain/model/user"
)

func main() {
	uri := fmt.Sprintf("postgres://%s/%s?sslmode=disable&user=%s&password=%s&port=%s&timezone=Asia/Tokyo",
		os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"))
	db, err := sql.Open("postgres", uri)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	log.Println("successfully connected to database")

	userFactory, err := user.NewUserFactory()
	if err != nil {
		panic(err)
	}
	userRepository, err := user.NewUserRepository(db)
	if err != nil {
		panic(err)
	}
	userService, err := user.NewUserService(userRepository)
	if err != nil {
		panic(err)
	}

	userApplicationService, err := user.NewUserApplicationService(userFactory, userRepository, *userService)
	if err != nil {
		panic(err)
	}

	if err := userApplicationService.Register("test-user"); err != nil {
		log.Fatal(err)
	}
}
