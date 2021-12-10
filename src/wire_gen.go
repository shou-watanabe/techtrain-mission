// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/jinzhu/gorm"
	"techtrain-mission/src/infra/repository"
	"techtrain-mission/src/presen/handler"
	"techtrain-mission/src/usecase"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func initUserHandler(db *gorm.DB) handler.UserHandler {
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)
	return userHandler
}
