package database

import (
	"log"
	"os"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=" + os.Getenv("HOST") +
		" user=" + os.Getenv("USER") +
		" password=" + os.Getenv("PASSWORD") +
		" dbname=" + os.Getenv("DBNAME") +
		" port=" + os.Getenv("PORT") +
		" sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panicf("Erro ao conectar com o banco de dados: %v", err)
	}

	err = DB.AutoMigrate(&models.Aluno{})
	if err != nil {
		log.Panicf("Erro ao realizar as migrações do banco de dados: %v", err)
	}
}
