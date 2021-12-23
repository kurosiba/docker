package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func GetSkillsHandler(w http.ResponseWriter, r *http.Request) {

	type PersonData struct {
		Name  string `json:"name"`
		Token int    `json:"token"`
	}

	var jsonData []PersonData
	var personData PersonData

	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/test_db")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT username, token FROM users")

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {

		if err := rows.Scan(&personData.Name, &personData.Token); err != nil {
			fmt.Println(err)
		}

		jsonData = append(jsonData, personData)
	}

	j, err := json.Marshal(jsonData)

	if err != nil {
		fmt.Println("error:", err)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(j))
	fmt.Fprint(w, string(j))
}

func main() {
	http.HandleFunc("/getAll", GetSkillsHandler)
	http.ListenAndServe(":8080", nil)
}
