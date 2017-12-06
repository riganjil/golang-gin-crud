package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

//Person struktur
type Person struct {
	ID           uint   `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Organization string `json:"organization"`
	City         string `json:"city"`
}

func main() {
	db, _ = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.SingularTable(true)
	db.AutoMigrate(&Person{})
	r := gin.Default()
	r.GET("/person/", GetPerson)
	r.GET("/person/:id", GetPeople)
	r.POST("/person", CreatePerson)
	r.PUT("/person/:id", UpdatePerson)
	r.DELETE("/person/:id", DeletePerson)
	r.Run(":3001")
}

//DeletePerson function for deleting person
func DeletePerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person
	d := db.Where("id = ?", id).Delete(&person)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

//UpdatePerson function for updating person
func UpdatePerson(c *gin.Context) {
	var person Person
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&person)
	db.Save(&person)
	c.JSON(200, person)
}

//CreatePerson function for creating person
func CreatePerson(c *gin.Context) {
	var person Person
	c.BindJSON(&person)
	db.Create(&person)
	c.JSON(200, person)
}

//GetPeople function for get one person
func GetPeople(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, person)
	}
}

// GetPerson function for getting list
func GetPerson(c *gin.Context) {
	var person []Person
	if err := db.Find(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, person)
	}
}
