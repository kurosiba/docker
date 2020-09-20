
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

func encoding(w http.ResponseWriter, name string, token int) {

  type PersonData struct{
    Name string
    Token int
  }

  data := PersonData{
    Name: name,
    Token: token,
  }

  j, err := json.Marshal(data)

  if err != nil{
    fmt.Println("error:", err)
  }

  log.Println(string(j))
  fmt.Fprint(w,string(j))
  
  //return 

}

func GetHandler(w http.ResponseWriter, r *http.Request){
  
  db, err := sql.Open("mysql", "root:root@tcp(mysql_test:3306)/test_db")

  if err != nil {
   panic(err)
  }

  defer db.Close()

  var (
    name string 
    token int 
  )

  if err := db.QueryRow("SELECT username,token FROM users WHERE token=?", 2).Scan(&name,&token);

  err != nil {
   log.Fatal(err)
  }

  //log.Println(name, token)
  
  encoding(w, name, token)
  
}

func AllGetHandler(w http.ResponseWriter, r *http.Request){

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
   
   var (
     token int
     name string
   )

  if err := rows.Scan(&name,&token);
  
  err != nil{
    log.Fatal(err)
  }

  //log.Println(name, token)
  encoding(w, name, token)

  }

  if err := rows.Err();

  err != nil{
   log.Fatal(err)
  }

}

func main(){
  http.HandleFunc("/", HelloHandler)
  http.HandleFunc("/get", GetHandler)
  http.HandleFunc("/getAll",AllGetHandler)
  /*when enter "http://localhost:8080/get" url, it send GerRequest to server*/
  //http.HandleFunc("/post", PostHandler)
  //http.HandleFunc("/put", PutHandler)

  http.ListenAndServe(":8080", nil)
  
}






