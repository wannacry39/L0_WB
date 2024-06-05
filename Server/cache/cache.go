package cache

import (
	"database/sql"
	"log"
)

var cache map[string]string = map[string]string{}

type order struct {
	id       string
	data_str string
}

func Cache_init() {
	connstr := "user=user_l0 password=qwerty host=localhost dbname=l0 sslmode=disable"
	db, err := sql.Open("postgres", connstr)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select data->>'order_uid' id, data from \"order\";")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		o := order{}
		err := rows.Scan(&o.id, &o.data_str)
		if err != nil {
			log.Fatal(err)
		}
		cache[o.id] = o.data_str
	}
}

func Add_data(id string, data string) {
	cache[id] = data
}

func Exists(id string) bool {
	_, ok := cache[id]
	return ok
}

func GetValue(id string) string {
	return cache[id]
}
