package initializers

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)
var DBM *gorm.DB;

func ConnectToMyDB(){
	dsn := "root:P9g78l@5b@/mydb"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DBM = db
}