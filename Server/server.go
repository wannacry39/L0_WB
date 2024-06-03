package Server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func Server() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "Server/template/index.html")
	})

	http.HandleFunc("/getform", func(w http.ResponseWriter, r *http.Request) {
		var data string

		connstr := "user=user_l0 password=qwerty host=localhost dbname=l0 sslmode=disable"
		db, err := sql.Open("postgres", connstr)

		if err != nil {
			log.Fatal(err)
		}

		defer db.Close()

		id := r.FormValue("order_id")
		row, err := db.Query("select data from \"order\" where data->>'order_uid'=$1;", id)

		if err != nil {
			log.Fatal(err)
		}
		defer row.Close()

		for row.Next() {
			err := row.Scan(&data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprint(w, data)
		}

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
