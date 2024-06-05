package Server

import (
	"database/sql"
	"fmt"
	"l0_project/Server/cache"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func Server() {
	cache.Cache_init()
	fmt.Println("Cache initialized")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "Server/template/index.html")
	})

	http.HandleFunc("/getform", func(w http.ResponseWriter, r *http.Request) {
		var data string
		id := r.FormValue("order_id")

		if cache.Exists(id) {
			fmt.Fprint(w, cache.GetValue(id))
		} else {
			connstr := "user=user_l0 password=qwerty host=localhost dbname=l0 sslmode=disable"
			db, err := sql.Open("postgres", connstr)

			if err != nil {
				log.Fatal(err)
			}

			defer db.Close()

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
			cache.Add_data(id, data)
		}
	})
	fmt.Println("Server is listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
