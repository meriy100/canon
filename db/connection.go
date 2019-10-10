package db

import (
	"github.com/go-yaml/yaml"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/meriy100/canon/models"
	"github.com/qor/validations"
	"io/ioutil"
	"os"
)

type ConnectionInfo struct {
	Dbms string
	User string
	Pass string
	Host string
	Port string
	Dbname string
	Sslmode string
}

func databaseConnectionInfo() (ConnectionInfo, error) {
	buf, err := ioutil.ReadFile("./db/database.yml")
	if err != nil {
		return ConnectionInfo{
			"postgres",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_SSL_MODE"),
			}, nil
	}

	data, err := ReadOnSliceMap(buf)
	return data, err
}


func ReadOnSliceMap(fileBuffer []byte) (ConnectionInfo, error) {
	data := make(map[string]ConnectionInfo, 20)
	err := yaml.Unmarshal(fileBuffer, &data)
	return data["development"], err
}

func GormConnect() *gorm.DB {
	connectionInfo, err := databaseConnectionInfo()
	if err != nil {
		panic(err)
	}
	//CONNECT := "host=" + HOST + " port=" + PORT + " user=" + USER + "password=" + PASS + " dbname=" + DBNAME
	CONNECT := "host=" + connectionInfo.Host +
		" user=" + connectionInfo.User +
		" password=" + connectionInfo.Pass +
		" dbname=" + connectionInfo.Dbname +
		" port=" + connectionInfo.Port +
		" sslmode=" + connectionInfo.Sslmode
	db,err := gorm.Open(connectionInfo.Dbms, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	validations.RegisterCallbacks(db)

	return db
}

func DropTables() {
	db := GormConnect()
	defer db.Close()

	db.DropTableIfExists("users")
}



func Migration() {
	db := GormConnect()
	defer db.Close()

	db.AutoMigrate(&models.User{})
}
