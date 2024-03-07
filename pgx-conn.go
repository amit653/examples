package main
import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func print_csv() {

	file, err := os.Open("dummy.csv")
	//*os.File

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", file)
	//fmt.Printf("%v",file.Stat())
	defer file.Close()
	reader := csv.NewReader(file) // Parameter is interface which can accept any type , here its File  struct
	//io.Reader

	record, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for i, rec := range record {
		fmt.Println(i, rec)

	}

}
func print_single_row(conn *pgx.Conn) (string, string) {
	var a, b string
	err := conn.QueryRow(context.Background(), "select stock_name ,stock_price from stock.stock_price where stock_name=$1", "TCS").Scan(&a, &b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	//fmt.Println(a, b)
	return a, b
}
func main() {
	type stock_struct struct {
		stock_name  string
		stock_price float32
	}

	//print_csv()
	//export DATABASE_URL="postgres://amitg:postgres@localhost:5433/stockdb"
	fmt.Println(os.Getenv("DATABASE_URL"))
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	//fmt.Println(print_single_row(conn)) //  sinlge row call
	
	
	///////////// display all rows using pgx.ForEachRow /////////
	rows, err := conn.Query(context.Background(), "select stock_name ,stock_price from stock.stock_price ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var r stock_struct
	_, err = pgx.ForEachRow(rows, []any{&r.stock_name, &r.stock_price}, func() error {
		fmt.Printf("%v, %v\n", r.stock_name, r.stock_price)
		return nil
	})
	fmt.Println("row call")
	if err != nil {
		fmt.Printf("ForEachRow error: %v", err)
		return
	}
	//fmt.Println(reflect.ValueOf(rowArray).Kind(), reflect.TypeOf(rowArray), reflect.ValueOf(rowArray))
	//fmt.Println("\n")
	//fmt.Println(reflect.ValueOf(rowArray1).Kind(), reflect.TypeOf(rowArray1), reflect.ValueOf(rowArray1))
	//fmt.Println("\n")
	//fmt.Println(reflect.ValueOf(rows).Kind(), reflect.TypeOf(rows), reflect.ValueOf(rows))
	/*
		rows1, err := conn.Query(context.Background(), "select stock_name ,stock_price from test ")
		if err != nil {
			log.Fatal(err)
		}
		defer rows1.Close()
		var rowSlc []stock_struct
		for rows1.Next() {   //////////////// display all rows using rows.next /////////
			rows1.Scan(&r.stock_name, &r.stock_price)
			rowSlc = append(rowSlc, r)

		}

		for _, j := range rowSlc {
			fmt.Printf("%v- %v\n", j.stock_name, j.stock_price)

		}
	*/

}
