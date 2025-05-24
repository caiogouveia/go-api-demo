package repository

import (
	"database/sql"
	"fmt"
	"goapidemo/model"
)

type VeiculoRepository struct {
	connection *sql.DB
}

func NewVeiculoRepository(connection *sql.DB) VeiculoRepository {
	return VeiculoRepository{
		connection: connection,
	}
}

func (pr *VeiculoRepository) GetVeiculos() ([]model.Veiculo, error) {
	query := "SELECT * FROM veiculo"
	rows, err := pr.connection.Query(query)

	fmt.Println("Hello World!")
	fmt.Println(rows)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var veiculoList []model.Veiculo
	var veiculoObj model.Veiculo

	for rows.Next() {
		err = rows.Scan(
			&veiculoObj.ID,
			&veiculoObj.VEICULO_MARCA,
			&veiculoObj.VEICULO_PLACA)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		veiculoList = append(veiculoList, veiculoObj)
	}

	rows.Close()
	return veiculoList, nil
}

func (pr *VeiculoRepository) CreateVeiculo(veiculo model.Veiculo) (int, error) {

	var id int
	query, err := pr.connection.Prepare("INSERT INTO veiculo" +
		"(VEICULO_MARCA, VEICULO_PLACA) VALUES ($1, $2) RETURNING ID")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(veiculo.VEICULO_MARCA, veiculo.VEICULO_PLACA).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (pr *VeiculoRepository) GetVeiculoById(id_veiculo int) (*model.Veiculo, error) {
	/*query := "SELECT * FROM veiculo WHERE ID = $1"
	row := pr.connection.QueryRow(query, id_veiculo)

	var veiculo model.Veiculo
	err := row.Scan(
		&veiculo.ID,
		&veiculo.VEICULO_MARCA,
		&veiculo.VEICULO_PLACA)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &veiculo, nil*/
	query, err := pr.connection.Prepare("SELECT * FROM veiculo WHERE ID = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var veiculo model.Veiculo
	err = query.QueryRow(id_veiculo).Scan(
		&veiculo.ID,
		&veiculo.VEICULO_MARCA,
		&veiculo.VEICULO_PLACA)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &veiculo, nil

}

/*
func (pr *VeiculoRepository) GetVeiculoByIdv2(id_veiculo int) (*model.Veiculo, error) {
	query, err := pr.connection.Prepare("SELECT * FROM veiculo WHERE ID = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var veiculo model.Veiculo
	err = query.QueryRow(id_veiculo).Scan(
		&veiculo.ID,
		&veiculo.VEICULO_MARCA,
		&veiculo.VEICULO_PLACA)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &veiculo, nil
}
*/
