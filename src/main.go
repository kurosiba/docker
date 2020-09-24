
package main

import (
	"fmt"
	"net/http"
  "log"
  "encoding/json"
  "database/sql"
   _ "github.com/go-sql-driver/mysql"  
)

func HelloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, World")
}

func GetHandler(w http.ResponseWriter, r *http.Request){

  type PersonData struct{
    Name string `json:"name"`
    Token int `json:"token"`
  }
  
  var jsonData []PersonData
  var personData PersonData

  db, err := sql.Open("mysql", "root:root@tcp(mysql_test:3306)/test_db")

  if err != nil {
   panic(err)
  }

  defer db.Close()

  if err := db.QueryRow("SELECT username,token FROM users WHERE token=?", 2).Scan(&personData.Name,&personData.Token);

  err != nil {
   log.Fatal(err)
  }

  jsonData = append(jsonData, personData)
  
  j, err := json.Marshal(jsonData)

  if err != nil{
    fmt.Println("error:", err)
  }
 
  log.Println(string(j))
  fmt.Fprint(w,string(j)) 
}

func AllGetHandler(w http.ResponseWriter, r *http.Request){
   
  type PersonData struct{
    Name string `json:"name"`
    Token int `json:"token"`
  }

  var jsonData []PersonData
  var personData PersonData

  db, err := sql.Open("mysql", "root:root@tcp(mysql_test:3306)/test_db")
  
  if err != nil {
  panic(err)
  }

  defer db.Close()

  rows, err := db.Query("SELECT username, token FROM users")

  if err != nil{
  log.Fatal(err)
  }

  defer rows.Close()

  for rows.Next(){
   
    if err := rows.Scan(&personData.Name,&personData.Token);
  
    err != nil{
      log.Fatal(err)
    }

    jsonData = append(jsonData, personData)
  }
  
  j, err := json.Marshal(jsonData)

  if err != nil{
    fmt.Println("error:", err)
  }

  if err := rows.Err();

  err != nil{
   log.Fatal(err)
  }
  
  log.Println(string(j))
  fmt.Fprint(w,string(j))
}

func main(){
  http.HandleFunc("/", HelloHandler)
  http.HandleFunc("/get", GetHandler)
  http.HandleFunc("/getAll",AllGetHandler)  
  http.ListenAndServe(":8080", nil)
 }

          




