package models

import (
 "log"
 db "../database"
 _ "github.com/lib/pq"
)

type User struct {
 Id        int    `json:"uid" form:"uid"`
 UserName string `json:"username" form:"username"`
 Email  string `json:"email" form:"email"`
 CreatedDate  string `json:"createddate" form:"createddate"`
}

func (p *User) AddPerson() (id int64, err error) {

 stmt, err := db.Pdb.Prepare("INSERT INTO userinfo(username,email,createddate) VALUES($1,$2,$3) RETURNING uid")
 if err != nil {
  log.Fatal(err.Error())
  return
 }

 res, err := stmt.Exec(p.UserName, p.Email, p.CreatedDate)

 if err != nil {
  log.Fatal(err.Error())
  return
 }
 
 id, err = res.LastInsertId()
 return
}

func (p *User) GetPersons() (users []User, err error) {
 
 users = make([]User, 0)
 rows, err := db.Pdb.Query("SELECT uid, username,email,createddate FROM userinfo")
 defer rows.Close()

 if err != nil {
  log.Fatal(err.Error())
  return
 }

 for rows.Next() {
  var user User
  rows.Scan(&user.Id, &user.UserName, &user.Email, &user.CreatedDate)
  users = append(users, user)
 }
 
 if err = rows.Err(); err != nil {
  log.Fatal(err.Error())
  return
 }
 return
}
