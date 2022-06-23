package repository_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/internal/usuarios/domain"
	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/internal/usuarios/repository"
	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	input := []domain.Usuario{
		{
			Id:        1,
			Nome:      "A",
			Sobrenome: "bcdef",
			Email:     "generico@email.com",
			Idade:     50,
			Altura:    1.89,
			Ativo:     true,
			Data:      "07/06/2022",
		},
		{
			Id:        2,
			Nome:      "B",
			Sobrenome: "cdefg",
			Email:     "generico@email.com",
			Idade:     25,
			Altura:    1.70,
			Ativo:     true,
			Data:      "08/06/2022",
		},
	}
	dataJson, _ := json.Marshal(input)
	dbMock := store.Mock{
		Data: dataJson,
		Err:  nil,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}

	myRepo := repository.NewRepository(&storeStub)
	resp, _ := myRepo.GetAll()
	assert.Equal(t, input, resp)
}

func TestUpdateSobrenomeIdade(t *testing.T) {
	input := []domain.Usuario{
		{
			Id:        1,
			Nome:      "A",
			Sobrenome: "Before Update",
			Email:     "generico@email.com",
			Idade:     50,
			Altura:    1.89,
			Ativo:     true,
			Data:      "07/06/2022",
		},
		{
			Id:        2,
			Nome:      "B",
			Sobrenome: "cdefg",
			Email:     "generico@email.com",
			Idade:     25,
			Altura:    1.70,
			Ativo:     true,
			Data:      "08/06/2022",
		},
	}

	expectedResult := domain.Usuario{
		Id:        1,
		Nome:      "A",
		Sobrenome: "After Update",
		Email:     "generico@email.com",
		Idade:     50,
		Altura:    1.89,
		Ativo:     true,
		Data:      "07/06/2022",
	}

	dataJson, _ := json.Marshal(input)
	dbMock := store.Mock{
		Data:          dataJson,
		Err:           nil,
		ReadWasCalled: false,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}

	myRepo := repository.NewRepository(&storeStub)
	resp, err := myRepo.UpdateSobrenomeIdade(1, "After Update", 50)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	assert.Equal(t, expectedResult, resp)
	assert.True(t, dbMock.ReadWasCalled)
}
