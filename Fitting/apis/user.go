package apis

import (
 "net/http"
 "log"
 "fmt"
 "gopkg.in/gin-gonic/gin.v1"
 . "../models"
)

func IndexApi(c *gin.Context) {
 c.String(http.StatusOK, "It works")
}

func AddPersonApi(c *gin.Context) {
 userName := c.Request.FormValue("username")
 email := c.Request.FormValue("email")
 createdDate := c.Request.FormValue("createddate")
 
 p := User{UserName: userName, Email: email, CreatedDate: createdDate}

 ra, err := p.AddPerson()
 if err != nil {
  log.Fatalln(err)
 }
 
 msg := fmt.Sprintf("insert successful %d", ra)
 c.JSON(http.StatusOK, gin.H{
  "msg": msg,
 })
}

func GetPersonsApi(c *gin.Context) {

 p := User{}
 ra, err := p.GetPersons()
 if err != nil {
  log.Fatalln(err)
 }
 
 c.JSON(http.StatusOK, gin.H{
  "msg": ra,
 })
}