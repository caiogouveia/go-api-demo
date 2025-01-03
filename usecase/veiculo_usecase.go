package usecase

import (
	"goapidemo/model"
	"goapidemo/repository"
)

type VeiculoUsecase struct {
	repository repository.VeiculoRepository
}

func NewVeiculoUsecase(repo repository.VeiculoRepository) VeiculoUsecase {
	return VeiculoUsecase{
		repository: repo,
	}
}

func (pu *VeiculoUsecase) GetVeiculos() ([]model.Veiculo, error) {
	return pu.repository.GetVeiculos()
}

func (pu *VeiculoUsecase) CreateVeiculo(veiculo model.Veiculo) (model.Veiculo, error) {
	veiculoId, err := pu.repository.CreateVeiculo(veiculo)
	if err != nil {
		return model.Veiculo{}, err
	}
	veiculo.ID = veiculoId
	return veiculo, nil
}

func (pu *VeiculoUsecase) GetVeiculoById(idveiculo int) (*model.Veiculo, error) {
	veiculo, err := pu.repository.GetVeiculoById(idveiculo)

	if err != nil {
		return nil, err
	}

	return veiculo, nil
}
