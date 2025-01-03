create table veiculo(
	id SERIAL primary key,
	veiculo_marca VARCHAR(50) not null,
	veiculo_placa VARCHAR(50) not null
);



select * from veiculo

insert into veiculo(veiculo_marca, veiculo_placa) values ('SCANIA','AAA-3421')