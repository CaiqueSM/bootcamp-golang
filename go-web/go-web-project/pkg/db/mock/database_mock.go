package mock

import(
	"database/sql"
	"github.com/DATA-DOG/go-txdb"
	_"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func init(){
	txdb.Register("txdb", "mysql", "root:AM@U3c=62@tcp(localhost:3306)/storage")
}

func InitDb()(*sql.DB, error){
	db, err := sql.Open("txdb", uuid.New().String())

	if err == nil{
		return db, db.Ping()
	}

	return db, err
}