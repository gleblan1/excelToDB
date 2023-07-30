package domain

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
)

type MySQLRepository struct {
	db *sql.DB
}

// NewMySQLRepository создает новый экземпляр MySQLRepository и инициализирует подключение к базе данных MySQL.
func NewMySQLRepository() *MySQLRepository {
	dbUser := "root"
	dbPassword := viper.GetString("dbPass")
	dbName := "excelToDB"
	dbHost := "localhost"
	dbPort := "3306"

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName))
	if err != nil {
		log.Fatal(err)
	}

	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) SaveData(data []Data) error {
	// Подготовка SQL-запроса для вставки данных
	query := "INSERT INTO data (column1, column2) VALUES (?, ?)"

	// Запускаем транзакцию
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	// Подготавливаем и выполняем запрос для каждой структуры Data
	for _, d := range data {
		_, err := tx.Exec(query, d.Column1, d.Column2)
		if err != nil {
			// Если произошла ошибка, отменяем транзакцию и возвращаем ошибку
			tx.Rollback()
			return err
		}
	}

	// Если все успешно, фиксируем транзакцию
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// GetData получает данные из базы данных.
func (r *MySQLRepository) GetData() ([]Data, error) {
	// Подготовка SQL-запроса для получения данных
	query := "SELECT column1, column2 FROM data"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []Data
	for rows.Next() {
		var d Data
		err := rows.Scan(&d.Column1, &d.Column2)
		if err != nil {
			log.Println(err)
			continue
		}
		data = append(data, d)
	}

	return data, nil
}
