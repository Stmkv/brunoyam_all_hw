// @title Example API
// @version 1.0
// @description This is a sample server
// @host localhost:8080
// @BasePath /
package main

import (
	_ "temp-prj/cmd/docs"
	"temp-prj/internal/repository/memstorage"
	"temp-prj/internal/server"
	"temp-prj/internal/service"
)

func main() {
	repo := memstorage.New()
	usersService := service.New(repo)

	srv := server.New(":8080", usersService)

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
