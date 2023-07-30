package main
//Reference
//https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
//https://www.enterprisedb.com/postgres-tutorials/everything-you-need-know-about-postgres-stored-procedures-and-functions
import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "postgres"
	dbname   = "stockdb"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	/////////////
	//update stock.stock_price set stock_price=%s,lastupdatetime=%s where stock_name=
	sqlStatement := `update test1 set a=1`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	//CALL transaction_test1()")
	sqlStatement = `call stock.proc_refresh_mv_plhistory();` //`CALL transaction_test1()`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}

}
