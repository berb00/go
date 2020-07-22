package main

import (
	"golang/base"
	"golang/mysql"
	"golang/server"
)

func baseModule() {
	// base.interface
	// base.DefInterface()

	// array
	// base.DefineUnidimensionalArray()
	// base.DefinedMultidimensionedArray()

	// point
	base.TestNew()
}

func mysqlModule() {
	mysql.LinkSQL()
	// mysql.InsertRowDemo("里斯1", 23)
	// mysql.InsertRowDemo("里斯2", 23)
	// mysql.InsertRowDemo("里斯3", 23)
	// mysql.InsertRowDemo("里斯4", 23)
	// mysql.UpdateRowDemo(24, 4)
	// mysql.DeleteRowDemo(4)
	// mysql.QueryRowDemo(2)
	mysql.QueryMultiRowDemo()

}

func serverModule() {
	mysql.LinkSQL()
	server.HttpServer()
}

func main() {
	// baseModule()
	// mysqlModule()
	serverModule()
}
