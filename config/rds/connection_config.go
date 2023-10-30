package rds

import (
	"fmt"
	"os"
)

func RDSConfiguration() (string, string) {
	masterDBHost := os.Getenv("MASTER_DB_HOST")
	masterDBPort := os.Getenv("MASTER_DB_PORT")
	masterDBUser := os.Getenv("MASTER_DB_USER")
	masterDBPassword := os.Getenv("MASTER_DB_PASS")
	masterDBName := os.Getenv("MASTER_DB_NAME")

	slaveDBHost := os.Getenv("SLAVE_DB_HOST")
	slaveDBPort := os.Getenv("SLAVE_DB_PORT")
	slaveDBUser := os.Getenv("SLAVE_DB_USER")
	slaveDBPassword := os.Getenv("SLAVE_DB_PASS")
	slaveDBName := os.Getenv("SLAVE_DB_NAME")

	// master rdsConnection Connection
	masterDBConnection := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		masterDBUser, masterDBPassword, masterDBHost, masterDBPort, masterDBName,
	)

	// slave rdsConnection connection
	slaveDBConnection := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		slaveDBUser, slaveDBPassword, slaveDBHost, slaveDBPort, slaveDBName,
	)

	return masterDBConnection, slaveDBConnection
}
