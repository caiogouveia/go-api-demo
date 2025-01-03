package main

import (
	"goapidemo/controller"
	"goapidemo/db"
	"goapidemo/repository"
	"goapidemo/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	//in.SetMode(gin.ReleaseMode)
	server := gin.Default()
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	/*
		Controller -> UseCase -> Repository
	*/
	//Camada de Repository
	VeiculoRepository := repository.NewVeiculoRepository(dbConnection)

	//Camada de UseCase
	VeiculoUsecase := usecase.NewVeiculoUsecase(VeiculoRepository)

	//Camada de Controllers
	VeiculoController := controller.NewVeiculoController(VeiculoUsecase)

	//Endpoint Selftest
	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong"})

	})

	//Endpoints GET
	server.GET("/veiculos", VeiculoController.GetVeiculos)
	server.GET("/veiculo/:veiculoId", VeiculoController.GetVeiculoById)

	//Endpoints POST
	server.POST("/criarveiculo", VeiculoController.CreateVeiculo)

	//Run server

	server.Run(":8000")

}
