package ddl

import (
	"database/sql"
	"fmt"
	"gostock/config"
	"gostock/server"
	"os"
)

func createDatabase() error {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		config.Data.Mysql.Username,
		config.Data.Mysql.Password,
		config.Data.Mysql.Host,
		config.Data.Mysql.Port)
	db, err := sql.Open("mysql", dns)
	defer db.Close()
	if err != nil {
		return err
	}

	ddl := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", config.Data.Mysql.Dbname)

	// 创建数据库
	_, err = db.Exec(ddl)
	if err != nil {
		return err
	}
	server.Log.Info("create databases success")
	return nil
}

func createTables() error {
	for _, tablename := range config.Data.DDL.Tables {

		filename := fmt.Sprintf("%s%s.sql", config.Data.DDL.Path, tablename)
		sqlBytes, err := os.ReadFile(filename)
		if err != nil {
			return err
		}
		ddl := string(sqlBytes)
		_, err = server.MysqlClient.Exec(ddl)
		if err != nil {
			return err
		}
		server.Log.Info("create table success " + tablename)
	}
	return nil
}

func Create() {
	err := createDatabase()
	if err != nil {
		server.Log.Error("create databases fail")
		return
	}
	err = createTables()
	if err != nil {
		server.Log.Error("create tables fail")
		return
	}
}
