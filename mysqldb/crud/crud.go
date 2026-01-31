package crud

import (
	"database/sql"
	"log"
	"mysqldb/crud/data/data"
	"mysqldb/crud/data/mysqldb"
)

func Insert(data data.Coluns) {
	var db *sql.DB
	db = mysqldb.Connection()
	_, err := db.Exec("insert into market_prod(nomeProd,preco,valAno,codeBar) values(?,?,?,?)", data.NomeProd, data.Preco, data.ValAno, data.CodeBar)
	if err != nil {
		log.Fatalln(err)
	}
}
func SelectAll() []data.Coluns {
	var colunsSli []data.Coluns

	var db *sql.DB
	db = mysqldb.Connection()
	row, err := db.Query("select * from market_prod")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	for row.Next() {
		var colun data.Coluns
		if err := row.Scan(&colun.Id, &colun.NomeProd, &colun.Preco, &colun.ValAno, &colun.CodeBar); err != nil {
			log.Fatal(err)
		}
		colunsSli = append(colunsSli, colun)
	}
	return colunsSli
}
