package belajar_golang_gorm

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	config := viper.New()
	config.SetConfigFile(".env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	if err != nil {
		return nil
	}

	username := config.GetString("DATABASE_USERNAME")
	password := config.GetString("DATABASE_PASSWORD")
	host := config.GetString("DATABASE_HOST")
	dbName := config.GetString("DATABASE_NAME")
	dbPort := config.GetString("DATABASE_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, dbPort, dbName)

	dialect := mysql.Open(dsn)
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

var db = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
}

func TestExecuteSQL(t *testing.T) {
	err := db.Exec("insert into sample(id, name) values(?, ?)", "1", "Jhon").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values(?, ?)", "2", "Doe").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values(?, ?)", "3", "Vale").Error
	assert.Nil(t, err)

	err = db.Exec("insert into sample(id, name) values(?, ?)", "4", "Rossi").Error
	assert.Nil(t, err)
}

type Sample struct {
	Id   string
	Name string
}

func TestRawSQL(t *testing.T) {
	var sample Sample
	err := db.Raw("select id, name from sample where id=?", "1").Scan(&sample).Error
	assert.Nil(t, err)
	assert.Equal(t, "Jhon", sample.Name)

	var samples []Sample
	err = db.Raw("select id, name from sample").Scan(&samples).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(samples))
}

func TestSqlRow(t *testing.T) {
	rows, err := db.Raw("select id, name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var samples []Sample
	for rows.Next() {
		var id string
		var name string

		err := rows.Scan(&id, &name)
		assert.Nil(t, err)

		samples = append(samples, Sample{
			Id:   id,
			Name: name,
		})
	}
	assert.Equal(t, 4, len(samples))
}

func TestScanRow(t *testing.T) {
	rows, err := db.Raw("select id, name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var samples []Sample
	for rows.Next() {
		err := db.ScanRows(rows, &samples)
		assert.Nil(t, err)
	}
	assert.Equal(t, 4, len(samples))
}
