package controller

import (
	"goapidemo/model"
	"goapidemo/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type veiculoController struct {
	veiculoUsecase usecase.VeiculoUsecase
}

// Inicializa o controller de veiculo
func NewVeiculoController(usecase usecase.VeiculoUsecase) veiculoController {
	return veiculoController{
		veiculoUsecase: usecase,
	}
}

func (p *veiculoController) GetVeiculos(ctx *gin.Context) {
	veiculos, err := p.veiculoUsecase.GetVeiculos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao buscar veiculos"})

	}

	//http faz ref. ao pacote net/http
	ctx.JSON(http.StatusOK, veiculos)

}

func (p *veiculoController) CreateVeiculo(ctx *gin.Context) {
	var veiculo model.Veiculo
	err := ctx.BindJSON(&veiculo)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Erro ao criar veiculo"})
		return
	}
	insertedVeiculo, err := p.veiculoUsecase.CreateVeiculo(veiculo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao criar veiculo"})
		return
	}
	ctx.JSON(http.StatusCreated, insertedVeiculo)
}

func (p *veiculoController) GetVeiculoById(ctx *gin.Context) {
	id := ctx.Param("veiculoId")

	if id == "" {
		response := model.Response{Message: "ID do veiculo não informado"} //gin.H{"message": "ID do veiculo não informado"})

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	idveiculo, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{Message: "ID do veiculo precisa ser um int"} //gin.H{"message": "ID do veiculo não informado"})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	veiculo, err := p.veiculoUsecase.GetVeiculoById(idveiculo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return

	}

	if veiculo == nil {
		response := model.Response{Message: "Veiculo não encontrado no DB"} //gin.H{"message": "ID do veiculo não informado"})
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	//http faz ref. ao pacote net/http
	ctx.JSON(http.StatusOK, veiculo)

}
