package main

import (
 "gopkg.in/gin-gonic/gin.v1"
 . "./apis"
)

func initRouter() *gin.Engine {
 router := gin.Default()

 router.GET("/", IndexApi)

 router.POST("/user", AddPersonApi)

 router.GET("/users", GetPersonsApi)

 //router.GET("/person/:id", GetPersonApi)

 //router.PUT("/person/:id", ModPersonApi)

 //router.DELETE("/person/:id", DelPersonApi)

 return router
}