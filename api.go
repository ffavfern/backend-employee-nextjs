package main

import (
	"example/go-api/model"
	_ "net/http"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Employee_id       string
	Employee_name     string
	Employee_username string
	Employee_password string
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/db_nextjs_employee?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	r := gin.Default()

	r.Use(cors.Default())

	//Find All
	r.GET("/employees", func(c *gin.Context) {
		var employee []model.Employee
		db.Find(&employee)
		c.JSON(200, employee)
	})

	//Find By Id
	r.GET("/employees/:id", func(c *gin.Context) {
		id := c.Param("id")
		var employee model.Employee
		db.First(&employee, id)
		c.JSON(200, employee)
	})

	//Post Data
	r.POST("/employees", func(c *gin.Context) {
		var employee model.Employee
		c.Bind(&employee)
		db.Create(&employee)
		c.JSON(200, gin.H{"data": employee})
	})

	//Delete Data
	r.DELETE("/employees/:id", func(c *gin.Context) {
		id := c.Param("id")
		var employee model.Employee
		db.First(&employee, id)
		db.Delete(&employee)
		c.JSON(200, employee)
	})

	//Put Data
	r.PUT("/employees/:id", func(c *gin.Context) {

		c.Header("Content-Type", "application/json")

		id := c.Param("id")

		var employee model.Employee
		var updatedemployee model.Employee

		c.BindJSON(&employee)
		c.BindJSON(&updatedemployee)

		db.First(&updatedemployee, id)

		updatedemployee.Employee_id = employee.Employee_id
		updatedemployee.Employee_name = employee.Employee_name
		updatedemployee.Employee_username = employee.Employee_username
		updatedemployee.Employee_password = employee.Employee_password

		db.Save(updatedemployee)
		c.JSON(200, updatedemployee)

		//c.JSON(200, gin.H{"id : " + id: "is puted"});
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
