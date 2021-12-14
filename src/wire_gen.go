// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"techtrain-mission/src/infra/repository"
	"techtrain-mission/src/presen/handler"
	"techtrain-mission/src/usecase"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func initUserHandler(driver *sql.DB) handler.UserHandler {
	userRepository := repository.NewUserRepository(driver)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)
	return userHandler
}

func initCharaHandler(driver *sql.DB) handler.CharaHandler {
	userRepository := repository.NewUserRepository(driver)
	userCharaRepository := repository.NewUserCharaRepository(driver)
	charaUsecase := usecase.NewCharaUsecase(userRepository, userCharaRepository)
	charaHandler := handler.NewCharaHandler(charaUsecase)
	return charaHandler
}

func initGachaHandler(driver *sql.DB) handler.GachaHandler {
	userRepository := repository.NewUserRepository(driver)
	charaRepository := repository.NewCharaRepository(driver)
	userCharaRepository := repository.NewUserCharaRepository(driver)
	gachaUsecase := usecase.NewGachaUsecase(userRepository, charaRepository, userCharaRepository)
	gachaHandler := handler.NewGachaHandler(gachaUsecase)
	return gachaHandler
}
