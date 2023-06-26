CREATE TABLE IF NOT EXISTS pacientes (
	id integer primary KEY NOT NULL,
	prioridade CHAR(1) NULL,
	paciente VARCHAR(48) NOT NULL,
	sexo CHAR(1) NULL,
	id_paciente VARCHAR(100) NOT NULL,
	procedimento VARCHAR(48) NOT NULL,
	modalidade VARCHAR(48) NOT NULL,
	data_exame DATE NOT NULL,
	data_laudo DATE NOT NULL,
	
	CHECK (sexo in ('F', 'M'))
);

CREATE TABLE IF NOT EXISTS users (
	id SERIAL NOT NULL,
	name VARCHAR(100) NOT NULL,
	email VARCHAR(100) NOT NULL,
	password VARCHAR(200) NOT NULL,
	created_at DATE,
	updated_at DATE,
	deleted_at DATE,
		
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS especialidades (
	id SERIAL NOT NULL,
	especialidade VARCHAR(100) NOT NULL,
	created_at DATE,
	updated_at DATE,
	deleted_at DATE,
		
	PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS medicos (
	id SERIAL NOT NULL,
	nome VARCHAR(100) NOT NULL,
	sexo CHAR(1) NULL,
	data_nasc DATE,
	cpf VARCHAR(11) NOT NULL,
	especialidade VARCHAR(100) NOT NULL,
	rqe VARCHAR(100),
	crm_estado VARCHAR(100),
	crm_numero VARCHAR(100),
	end_logradouro VARCHAR(100),
	end_numero VARCHAR(100),
	end_complemento VARCHAR(100),
	end_bairro VARCHAR(100),
	end_cidade VARCHAR(100),
	end_uf VARCHAR(100),
	end_cep VARCHAR(100),
	contato_email VARCHAR(100),
	contato_telefone1 VARCHAR(100),
	contato_telefone2 VARCHAR(100),
	contato_telefone3 VARCHAR(100),
	contato_obs VARCHAR(100),
	created_at DATE,
	updated_at DATE,
	deleted_at DATE,
		
	PRIMARY KEY (id),
	CHECK (sexo in ('F', 'M'))
);