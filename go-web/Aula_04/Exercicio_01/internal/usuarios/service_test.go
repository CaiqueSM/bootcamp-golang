package usuarios_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/internal/usuarios"
	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/internal/usuarios/domain"
	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
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
		Email:     "alternativo@email.com",
		Idade:     49,
		Altura:    1.90,
		Ativo:     true,
		Data:      "08/06/2022",
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

	myRepo := usuarios.NewRepository(&storeStub)
	myService := usuarios.NewService(myRepo)
	resp, err := myService.Update(
		1, "A","After Update","alternativo@email.com", 49, 1.90, true, "08/06/2022")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	assert.Equal(t, expectedResult, resp)
	assert.True(t, dbMock.ReadWasCalled)
}

func TestDelete(t *testing.T) {
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

	expectedResult := []domain.Usuario{
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

	myRepo := usuarios.NewRepository(&storeStub)
	myService := usuarios.NewService(myRepo)
	err1 := myService.Delete(2)
	resp,_ := myService.GetAll()

	assert.Equal(t, expectedResult, resp)
	assert.Nil(t, err1)
}

func TestDeleteError(t *testing.T) {
	expectedError := errors.New("error for delete")
	dbMock := store.Mock{
		Err: expectedError,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := usuarios.NewRepository(&storeStub)
	myService := usuarios.NewService(myRepo)

	err := myService.Delete(3)

	assert.Equal(t, expectedError, err)
}