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

INSERT INTO pacientes
    VALUES (DEFAULT, '1', 'Paulo Ricardo', 'M', '3123489', 'TC Crânio', 'CT', '01/20/2023', '01/20/2023'),
	   (DEFAULT, '1', 'João Rodrigues de Almeida', 'M', '3246600', 'TORAX', 'CR', '01/03/2023', '01/03/2023'),
	   (DEFAULT, '1', 'Karla de Souza', 'F', '3369711', 'ABDOME TOTAL', 'US', '02/24/2023', '02/24/2023'),
	   (DEFAULT, '1', 'Diógenes Batista dos Santos', 'M', '3492822', 'ABDOME TOTAL', 'CT', '01/07/2023', '01/07/2023'),
	   (DEFAULT, '1', 'Juliano Righetto', 'M', '3615933', 'HEAD', 'CT', '04/19/2023', '04/19/2023'),
	   (DEFAULT, '1', 'Francisco Silva', 'M', '3739044', 'TORNOZELO', 'CT', '05/11/2023', '05/11/2023'),
	   (DEFAULT, '1', 'Mariana dos Santos Neves', 'F', '3862155', 'MAMOGRAFIA', 'CR', '04/09/2023', '04/09/2023'),
	   (DEFAULT, '1', 'Juliana Battisti', 'F', '3874466', 'BACIA', 'CT', '03/18/2023', '03/18/2023'),
	   (DEFAULT, '1', 'Luiza Patrícia dos Santos', 'F', '3886777', 'TC Crânio', 'CT', '03/10/2023', '03/10/2023'),
	   (DEFAULT, '1', 'Maria Eduarda Félix de Almeida', 'F', '3899088', 'MAMOGRAFIA', 'CR', '04/17/2023', '04/17/2023'),
	   (DEFAULT, '1', 'Patrícia Rouanet', 'F', '3911399', 'PESCOCO', 'CT', '02/23/2023', '02/23/2023'),
	   (DEFAULT, '1', 'Mary de Souza', 'F', '3923710', 'ABDOME TOTAL', 'CT', '01/10/2023', '01/10/2023'),
	   (DEFAULT, '1', 'João de Souza e Silva', 'M', '3936021', 'COLUNAS', 'CR', '06/18/2023', '06/18/2023'),
	   (DEFAULT, '1', 'Luciano Patrício', 'M', '3948332', 'TC Crânio', 'CT', '02/21/2023', '02/21/2023'),
	   (DEFAULT, '1', 'Jaqueline Belgrado', 'F', '3960643', 'MAMOGRAFIA', 'CR', '02/12/2023', '02/12/2023'),
	   (DEFAULT, '1', 'Arcênio Segundo', 'M', '3972954', 'COLUNA LOMBAR', 'MR', '03/11/2023', '03/11/2023'),
	   (DEFAULT, '1', 'Diniz Medeiros', 'M', '3985265', 'TC Crânio', 'CT', '01/21/2023', '01/21/2023'),
	   (DEFAULT, '1', 'Diogo Silva', 'M', '3997576', 'MUSCULO', 'MR', '01/31/2023', '01/31/2023'),
	   (DEFAULT, '1', 'Jackie Santini', 'F', '4009887', 'MAMOGRAFIA', 'CR', '03/29/2023', '03/29/2023'),
	   (DEFAULT, '1', 'Rafael Medeiros', 'M', '4022198', 'COLUNA LOMBAR', 'MR', '05/28/2023', '05/28/2023');

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