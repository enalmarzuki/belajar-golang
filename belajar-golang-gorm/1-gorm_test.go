package belajar_golang_gorm

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"belajar-golang-gorm/model"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
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
	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

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

func TestCreateUser(t *testing.T) {
	user := model.User{
		ID:       "1",
		Password: "Rahasia",
		Name: model.Name{
			FirstName:  "Jhon",
			MiddleName: "Doe",
			LastName:   "Rossi",
		},
		Information: "ini akan di ignore",
	}

	response := db.Create(&user)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(1), response.RowsAffected)
}

func TestMultipleInsert(t *testing.T) {
	var users []model.User
	for i := 2; i < 10; i++ {
		users = append(users, model.User{
			ID:       strconv.Itoa(i),
			Password: "rahasia",
			Name: model.Name{
				FirstName: "User " + strconv.Itoa(i),
			},
		})
	}

	result := db.Create(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 8, int(result.RowsAffected))
}

func TestBatchInsert(t *testing.T) {
	var users []model.User
	for i := 10; i < 20; i++ {
		users = append(users, model.User{
			ID:       strconv.Itoa(i),
			Password: "rahasia",
			Name: model.Name{
				FirstName: "User Batch " + strconv.Itoa(i),
			},
		})
	}

	result := db.CreateInBatches(&users, 2)
	assert.Nil(t, result.Error)
	assert.Equal(t, 10, int(result.RowsAffected))
}

func TestTransactionSuccess(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&model.User{ID: "10", Password: "rahasia", Name: model.Name{FirstName: "User 10"}}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&model.User{ID: "11", Password: "rahasia", Name: model.Name{FirstName: "User 11"}}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&model.User{ID: "12", Password: "rahasia", Name: model.Name{FirstName: "User 12"}}).Error
		if err != nil {
			return err
		}

		return nil
	})

	assert.Nil(t, err)
}

func TestTransactionRollback(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&model.User{ID: "13", Password: "rahasia", Name: model.Name{FirstName: "User 13"}}).Error
		if err != nil {
			return err
		}

		err = tx.Create(&model.User{ID: "11", Password: "rahasia", Name: model.Name{FirstName: "User 11"}}).Error
		if err != nil {
			return err
		}

		return nil
	})

	assert.NotNil(t, err)
}

func TestManualTransactionSuccess(t *testing.T) {
	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Create(&model.User{ID: "13", Password: "rahasia", Name: model.Name{FirstName: "User 13"}}).Error
	assert.Nil(t, err)

	err = tx.Create(&model.User{ID: "14", Password: "rahasia", Name: model.Name{FirstName: "User 14"}}).Error
	assert.Nil(t, err)

	if err == nil {
		tx.Commit()
	}
}

func TestManualTransactionRollback(t *testing.T) {
	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Create(&model.User{ID: "15", Password: "rahasia", Name: model.Name{FirstName: "User 15"}}).Error
	assert.Nil(t, err)

	err = tx.Create(&model.User{ID: "14", Password: "rahasia", Name: model.Name{FirstName: "User 14"}}).Error
	assert.Nil(t, err)

	if err == nil {
		tx.Commit()
	}
}

func TestQuerySingleObject(t *testing.T) {
	user := model.User{}

	err := db.First(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, "1", user.ID)

	user = model.User{}
	err = db.Last(&user).Error
	assert.Nil(t, err)
	assert.Equal(t, "9", user.ID)

	user = model.User{}
	db.Take(&user)
	fmt.Println("USER ", user)
}

func TestQuerySingleObjectInlineCondition(t *testing.T) {
	user := model.User{}
	err := db.First(&user, "id = ?", "5").Error
	assert.Nil(t, err)
	assert.Equal(t, "5", user.ID)
	assert.Equal(t, "User 5", user.Name.FirstName)

	user = model.User{}
	err = db.Take(&user, "id = ?", "5").Error
	assert.Nil(t, err)
	assert.Equal(t, "5", user.ID)
	assert.Equal(t, "User 5", user.Name.FirstName)
}

func TestQueryAllObject(t *testing.T) {
	var users []model.User
	err := db.Find(&users, "id in ?", []string{"1", "2", "3", "4"}).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))
}

func TestQueryCondition(t *testing.T) {
	var users []model.User

	// untuk kondisi yang where().where() itu hasilnya akan menggunakan "AND"
	// contoh hasil query -> SELECT * FROM `users` WHERE first_name like '%User%' AND password = 'rahasia'
	err := db.Where("first_name like ?", "%User%").Where("password = ?", "rahasia").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 13, len(users))
}

func TestOrCondition(t *testing.T) {
	var users []model.User

	// untuk kondisi yang where().Or() itu hasilnya akan menggunakan "AND"
	// contoh hasil query -> SELECT * FROM `users` WHERE first_name like '%User%' OR password = 'rahasia'
	err := db.Where("first_name like ?", "%User%").Or("password = ?", "rahasia").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 14, len(users))
}

func TestNotCondition(t *testing.T) {
	var users []model.User

	// contoh hasil query -> SELECT * FROM `users` WHERE NOT first_name like '%User%' AND password = 'rahasia'
	err := db.Not("first_name like ?", "%User%").Where("password = ?", "rahasia").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}

func TestSelectFields(t *testing.T) {
	var users []model.User

	err := db.Select("id", "first_name").Find(&users).Error
	assert.Nil(t, err)

	for _, user := range users {
		assert.NotNil(t, user.ID)
		assert.NotEqual(t, "", user.Name.FirstName)
	}

	assert.Equal(t, 14, len(users))
}

func TestStructCondition(t *testing.T) {
	userCondition := model.User{
		Name: model.Name{
			FirstName: "User 5",
			LastName:  "", // tidak bisa karena dianggap default value
		},
		Password: "rahasia",
	}

	var users []model.User

	// contoh hasil query -> SELECT * FROM `users` WHERE `users`.`password` = 'rahasia' AND `users`.`first_name` = 'User 5'
	// field LastName tidak ada di query
	err := db.Where(userCondition).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}

func TestMapCondition(t *testing.T) {
	mapCondition := map[string]interface{}{
		"middle_name": "",
		"last_name":   "",
	}

	var users []model.User

	// contoh hasil query -> SELECT * FROM `users` WHERE `last_name` = '' AND `middle_name` = ''
	// field middle_name dan last_name ada di query
	err := db.Where(mapCondition).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 13, len(users))
}

func TestOrderLimitOffset(t *testing.T) {
	var users []model.User
	err := db.Order("id asc, first_name desc").Limit(5).Offset(5).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 5, len(users))
}

type UserResponse struct {
	ID        string
	FirstName string
	LastName  string
}

func TestQueryNonModel(t *testing.T) {
	var users []UserResponse
	err := db.Model(&model.User{}).Select("id, first_name, last_name").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 14, len(users))

	fmt.Println(users)
}

func TestUpdate(t *testing.T) {
	user := model.User{}
	err := db.Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)

	user.Name.FirstName = "Valantino"
	user.Name.MiddleName = ""
	user.Name.LastName = "Rossi"
	user.Password = "rahasia123"

	err = db.Save(&user).Error
	assert.Nil(t, err)
}

func TestUpdateSelectedColumns(t *testing.T) {
	err := db.Model(&model.User{}).Where("id = ?", "1").Updates(map[string]interface{}{
		"middle_name": "",
		"last_name":   "Morro",
	}).Error
	assert.Nil(t, err)

	err = db.Model(&model.User{}).Where("id = ?", "1").Update("password", "diubahlagi").Error
	assert.Nil(t, err)

	err = db.Where("id = ?", "1").Updates(model.User{
		Name: model.Name{
			FirstName: "Jhon",
			LastName:  "Doe",
		},
	}).Error
	assert.Nil(t, err)
}

func TestAutoIncrement(t *testing.T) {
	for i := 0; i < 10; i++ {
		userLog := model.UserLog{
			UserId: "1",
			Action: "Test Action",
		}

		err := db.Create(&userLog).Error
		assert.Nil(t, err)
		assert.NotEqual(t, 0, userLog.ID)
		fmt.Println(userLog.ID)
	}
}

func TestSaveOrUpdate(t *testing.T) {
	/*
		* Upsert (Update Insert)
		- Sebelumnya, kita telah menggunakan method Save() untuk melakukan UPDATE
		- Method Save() sebenarnya memiliki kemampuan untuk mendeteksi apakah harus melakukan Update atau Create
		- Jika data yang kita kirim tidak memiliki value ID, maka secara default akan melakukan Create
		- Jika data yang kita kirim memiliki value ID, maka akan melakukan Update
		- Hal ini mungkin cocok untuk jenis data yang ID nya adalah Auto Increment, karena kita tidak butuh ID ketika melakukan Create
	*/

	userLog := model.UserLog{
		UserId: "1",
		Action: "Test Action",
	}

	err := db.Save(&userLog).Error // Create
	assert.Nil(t, err)

	userLog.UserId = "2"
	err = db.Save(&userLog).Error // Update
	assert.Nil(t, err)
}

func TestSaveOrUpdateNonAutoIncrement(t *testing.T) {
	user := model.User{
		ID: "99",
		Name: model.Name{
			FirstName: "User 99",
		},
	}

	err := db.Save(&user).Error // Insert
	assert.Nil(t, err)

	user.Name.FirstName = "User 99 Updated"
	err = db.Save(&user).Error // Update
	assert.Nil(t, err)
}

func TestConflict(t *testing.T) {
	user := model.User{
		ID: "88",
		Name: model.Name{
			FirstName: "User 88",
		},
	}

	err := db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&user).Error // Insert
	assert.Nil(t, err)
}

func TestDelete(t *testing.T) {
	// Cara 1
	var user model.User
	err := db.Take(&user, "id = ?", "88").Error
	assert.Nil(t, err)

	err = db.Delete(&user).Error
	assert.Nil(t, err)

	// Cara 2
	err = db.Delete(&model.User{}, "id = ?", "99").Error
	assert.Nil(t, err)

	// Cara 3
	err = db.Where("id = ?", "77").Delete(&model.User{}).Error
	assert.Nil(t, err)
}

func TestSoftDelete(t *testing.T) {
	todo := model.Todo{
		UserId:      "1",
		Title:       "Todo 1",
		Description: "Description 1",
	}
	err := db.Create(&todo).Error
	assert.Nil(t, err)

	err = db.Delete(&todo).Error
	assert.Nil(t, err)
	assert.NotNil(t, todo.DeletedAt)

	var todos []model.Todo
	err = db.Find(&todos).Error
	assert.Nil(t, err)
	assert.Equal(t, 0, len(todos))
}

func TestUnscoped(t *testing.T) {
	var todo model.Todo
	/*
		error krn qeuery nya seperti ini
		SELECT * FROM `todos` WHERE id = 1 AND `todos`.`deleted_at` IS NULL ORDER BY `todos`.`id` LIMIT 1

		masih mencari row yang deleted_at nya === NULL
	*/
	// err := db.First(&todo, "id = ?", 1).Error

	// Solusi
	err := db.Unscoped().First(&todo, "id = ?", 3).Error
	assert.Nil(t, err)
	fmt.Println(todo)

	/*
		fungsi dibwah ini masih menjalankan soft delete, hasil query nya seperti ini
		UPDATE `todos` SET `deleted_at`='2025-05-14 15:21:05.473' WHERE `todos`.`id` = 1 AND `todos`.`deleted_at` IS NULL
	*/
	// err = db.Delete(&todo).Error

	// Solusinya
	err = db.Unscoped().Delete(&todo).Error
	assert.Nil(t, err)

	var todos []model.Todo
	err = db.Unscoped().Find(&todos).Error
	assert.Nil(t, err)
}

func TestLock(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var user model.User
		err := tx.Clauses(clause.Locking{Strength: clause.LockingStrengthUpdate}).Take(&user, "id = ?", "1").Error
		if err != nil {
			return err
		}

		user.Name.FirstName = "Joko"
		user.Name.LastName = "Morro"
		err = tx.Save(&user).Error
		return err
	})
	assert.Nil(t, err)
}

func TestCreateWallet(t *testing.T) {
	wallet := model.Wallet{
		ID:      "1",
		UserId:  "1",
		Balance: 1000000,
	}

	err := db.Create(&wallet).Error
	assert.Nil(t, err)
}

func TestRetrieveRalation(t *testing.T) {
	/*
		- Preload(string) => ini cocok untuk relasi one to many (Has Many) atau many to many
	*/
	var user model.User
	err := db.Model(&model.User{}).Preload("Wallet").Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)

	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "1", user.Wallet.ID)
	// assert.NotEqual(t, "1", user.Wallet.ID) // => sebelum ditambah preload
}

func TestRetrieveRalationJoin(t *testing.T) {
	/*
		- Joins(string) => ini cocok untuk relasi one to one (hasOne)
	*/
	var user model.User
	err := db.Model(&model.User{}).Joins("Wallet").Take(&user, "users.id = ?", "1").Error
	assert.Nil(t, err)

	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "1", user.Wallet.ID)
}

func TestAutoCreateUpdate(t *testing.T) {
	user := model.User{
		ID:       "20",
		Password: "rahasia",
		Name: model.Name{
			FirstName: "User 20",
		},
		Wallet: model.Wallet{
			ID:      "20",
			UserId:  "20",
			Balance: 1000000,
		},
	}

	err := db.Create(&user).Error
	assert.Nil(t, err)
}

func TestSkipAutoCreateUpdate(t *testing.T) {
	user := model.User{
		ID:       "21",
		Password: "rahasia",
		Name: model.Name{
			FirstName: "User 21",
		},
		Wallet: model.Wallet{
			ID:      "21",
			UserId:  "21",
			Balance: 1000000,
		},
	}

	err := db.Omit(clause.Associations).Create(&user).Error
	assert.Nil(t, err)
}

func TestUserAndAddresses(t *testing.T) {
	user := model.User{
		ID:       "52",
		Password: "rahasia",
		Name: model.Name{
			FirstName: "User 52",
		},
		Wallet: model.Wallet{
			ID:      "52",
			UserId:  "52",
			Balance: 1000000,
		},
		Addresses: []model.Address{
			{
				UserId:  "52",
				Address: "Jalan A",
			},
			{
				UserId:  "52",
				Address: "Jalan B",
			},
		},
	}

	err := db.Create(&user).Error
	assert.Nil(t, err)
}

func TestPreloadJoinOneToMany(t *testing.T) {
	var users []model.User
	err := db.Model(&model.User{}).Preload("Addresses").Joins("Wallet").Find(&users).Error
	assert.Nil(t, err)
}

func TestTakePreloadJoinOneToMany(t *testing.T) {
	var user model.User
	err := db.Model(&model.User{}).Preload("Addresses").Joins("Wallet").Take(&user, "users.id = ?", "50").Error
	assert.Nil(t, err)
}

func TestBelongsTo(t *testing.T) {
	fmt.Println("Preload")
	var addresses []model.Address
	err := db.Model(&model.Address{}).Preload("User").Find(&addresses).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(addresses))

	fmt.Println("Joins")
	addresses = []model.Address{}
	err = db.Model(&model.Address{}).Joins("User").Find(&addresses).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(addresses))
}

func TestBelongsToWallet(t *testing.T) {
	fmt.Println("Preload")
	var wallets []model.Wallet
	err := db.Model(&model.Wallet{}).Preload("User").Find(&wallets).Error
	assert.Nil(t, err)

	fmt.Println("Joins")
	wallets = []model.Wallet{}
	err = db.Model(&model.Wallet{}).Joins("User").Find(&wallets).Error
	assert.Nil(t, err)
}

func TestCreateManyToMany(t *testing.T) {
	product := model.Product{
		ID:    "P001",
		Name:  "Contoh Product",
		Price: 1000000,
	}
	err := db.Create(&product).Error
	assert.Nil(t, err)

	err = db.Table("user_like_product").Create(map[string]interface{}{
		"user_id":    "1",
		"product_id": "P001",
	}).Error
	assert.Nil(t, err)

	err = db.Table("user_like_product").Create(map[string]interface{}{
		"user_id":    "2",
		"product_id": "P001",
	}).Error
	assert.Nil(t, err)
}

func TestPreloadManyToMany(t *testing.T) {
	var product model.Product
	err := db.Preload("LikedByUsers").Take(&product, "id = ?", "P001").Error
	assert.Nil(t, err)
	assert.Equal(t, 2, len(product.LikedByUsers))
}

func TestPreloadManyToManyUser(t *testing.T) {
	var user model.User
	err := db.Preload("LikeProducts").Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(user.LikeProducts))
}

func TestAssociationFind(t *testing.T) {
	var product model.Product
	err := db.Take(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	var users []model.User
	err = db.Model(&product).Where("users.first_name LIKE ?", "User%").Association("LikedByUsers").Find(&users)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}

func TestAssociationAppend(t *testing.T) {
	var user model.User
	err := db.Take(&user, "id = ?", "3").Error
	assert.Nil(t, err)

	var product model.Product
	err = db.Take(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Append(&user)
	assert.Nil(t, err)
}

func TestAssociationReplace(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var user model.User
		err := tx.Take(&user, "id = ?", "1").Error
		assert.Nil(t, err)

		wallet := model.Wallet{
			ID:      "01",
			UserId:  user.ID,
			Balance: 1000000,
		}
		err = tx.Model(&user).Association("Wallet").Replace(&wallet)
		return err
	})
	assert.Nil(t, err) // ada error krn pada saat membuat tabel wallets kita pasang constraint NOT NULL di kolom user_id
}

func TestAssociationDelete(t *testing.T) {
	var user model.User
	err := db.Take(&user, "id = ?", "3").Error
	assert.Nil(t, err)

	var product model.Product
	err = db.Take(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Delete(&user)
	assert.Nil(t, err)
}

func TestAssociationClear(t *testing.T) {
	var product model.Product
	err := db.Take(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Clear() // => Clear() untuk menghapus semua relasi ditabel sebelahnya artinya product P001 relasi LikedByUsers akan jadi kosong
	assert.Nil(t, err)
}

func TestPreloadingWithCondition(t *testing.T) {
	var user model.User
	err := db.Preload("Wallet", "balance > ?", 100).Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)

	fmt.Println("USER : ", user)
}

func TestPreloadingNested(t *testing.T) {
	var wallet model.Wallet
	err := db.Preload("User.Addresses").Take(&wallet, "id = ?", "52").Error
	assert.Nil(t, err)

	fmt.Println("WALLET : ", wallet)
	fmt.Println("wallet.User : ", wallet.User)
	fmt.Println("wallet.User.Addresses : ", wallet.User.Addresses)
}

func TestPreloadingAll(t *testing.T) {
	var user model.User
	err := db.Preload(clause.Associations).Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)
}

func TestJoinQuery(t *testing.T) {
	var users []model.User
	err := db.Joins("join wallets on wallets.user_id = users.id").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))

	users = []model.User{}
	err = db.Joins("Wallet").Find(&users).Error // left Join
	assert.Nil(t, err)
	assert.Equal(t, 17, len(users))
}

func TestJoinWithCondition(t *testing.T) {
	var users []model.User
	err := db.Joins("join wallets on wallets.user_id = users.id AND wallets.balance > ?", 500000).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))

	users = []model.User{}
	err = db.Joins("Wallet").Where("Wallet.balance > ?", 500000).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))
}

func TestCount(t *testing.T) {
	var count int64
	err := db.Model(&model.User{}).Joins("Wallet").Where("Wallet.balance > ?", 500000).Count(&count).Error
	assert.Nil(t, err)
	assert.Equal(t, int64(4), count)
}

type AggregationResult struct {
	TotalBalance int64
	MinBalance   int64
	MaxBalance   int64
	AvgBalance   float64
}

func TestAggregation(t *testing.T) {
	var result AggregationResult
	err := db.Model(&model.Wallet{}).Select("sum(balance) as total_balance", "min(balance) as min_balance", "max(balance) as max_balance", "avg(balance) as avg_balance").Take(&result).Error
	assert.Nil(t, err)
	assert.Equal(t, int64(4000000), result.TotalBalance)
	assert.Equal(t, int64(1000000), result.MinBalance)
	assert.Equal(t, int64(1000000), result.MaxBalance)
	assert.Equal(t, float64(1000000), result.AvgBalance)
}

func TestAggregationGroupByAndHaving(t *testing.T) {
	var result []AggregationResult
	err := db.Model(&model.Wallet{}).Select("sum(balance) as total_balance", "min(balance) as min_balance", "max(balance) as max_balance", "avg(balance) as avg_balance").Joins("User").
		Group("User.id").
		Having("sum(balance) > ?", 500000).
		Find(&result).Error

	assert.Nil(t, err)
	assert.Equal(t, 4, len(result))
}

func TestContext(t *testing.T) {
	ctx := context.Background()

	var users []model.User
	err := db.WithContext(ctx).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 17, len(users))
}

func BrokeWalletBalance(db *gorm.DB) *gorm.DB {
	return db.Where("balance = ?", 0)
}

func SultanWalletBalance(db *gorm.DB) *gorm.DB {
	return db.Where("balance > ?", 1000000)
}

func TestScopes(t *testing.T) {
	var wallets []model.Wallet
	err := db.Scopes(BrokeWalletBalance).Find(&wallets).Error
	assert.Nil(t, err)

	wallets = []model.Wallet{}
	err = db.Scopes(SultanWalletBalance).Find(&wallets).Error
	assert.Nil(t, err)
}

func TestMigrator(t *testing.T) {
	// tidak disarankan pakai ini untuk real case
	err := db.Migrator().AutoMigrate(&model.GuestBook{})
	assert.Nil(t, err)
}

func TestHook(t *testing.T) {
	user := model.User{
		Password: "rahasia",
		Name: model.Name{
			FirstName: "User 100",
		},
	}

	err := db.Create(&user).Error
	assert.Nil(t, err)
	assert.NotEqual(t, "", user.ID)

	fmt.Println(user.ID)
}
