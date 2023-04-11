package main

import (
	"net/http"

	"yrkesakademin/yagolangapi/data"

	"github.com/gin-gonic/gin"
)

func start(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func main() {
	data.InitDatabase()
	router := gin.Default()
	router.GET("/", start)
	router.Run("localhost:8080")

	//e := data.Employee{
	//	Id:   0,
	//	Age:  1,
	//	City: "Stockholm",
	//	Name: "Natalia",
	//}

	//if e.IsCool() {
	//	fmt.Printf("Name is cool:%\n", e.Name)
	//}else {
	//	fmt.Printf("Name:%\n", e.Name)
	//}

	//fmt.Println("Hello")
	//t := tabby.New()
	//t.AddHeader("Name", "Age", "City")
	//t.AddLine("Natalia", "35", "Malmo")
	//t.AddLine("Alice", "5", "Malmo")
	//t.Print()
}
