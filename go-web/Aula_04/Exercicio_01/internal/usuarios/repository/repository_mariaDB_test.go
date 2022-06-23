package repository_test

import (
	"log"
	"testing"

	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/internal/usuarios/domain"
	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/internal/usuarios/repository"
	"github.com/CaiqueSM/bootcamp-golang.git/go-web/Aula_04/Exercicio_01/pkg/db"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	usuario := domain.Usuario{
		Id:        1,
		Nome:      "ABC",
		Sobrenome: "DEFG",
		Email:     "generico@Email.com",
		Idade:     25,
		Altura:    1.68,
		Ativo:     true,
		Data:      "23-06-2022",
	}
	db.Init()
	myRepo := repository.NewRepositoryMariaDB(db.StorageDB)
	i,_:=  myRepo.LastID()
	usuarioResult, err := myRepo.Store(1, "ABC", "DEFG", "generico@Email.com", 25, 1.68, true, "23-06-2022")
	if err!= nil{
		log.Println(err)
	}
	assert.Equal(t, i, int64(0))
	assert.Equal(t, usuario, usuarioResult)
}