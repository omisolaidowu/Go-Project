package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Arraytest struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" bson:"_id,omitempty"`
	First_name string             `bson:"name,omitempty bson:"name,omitempty"`
	Last_name  string             `bson:"name,omitempty bson:"name,omitempty"`
	Username   string             `bson:"name,omitempty bson:"name,omitempty"`
	Email      string             `bson:"name,omitempty bson:"name,omitempty"`
	Password   float32            `bson:"salary,omitempty bson:"salary,omitempty"`
}

// var emp1 = Arraytest{Name: "Idowu", Salary: 12345.09}

// var client *mongo.Client

var DB = connectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("myGoappDB").Collection("myUsers")
	return collection
}

func postData(c *gin.Context) {
	var userCollection = GetCollection(DB, "myUsers")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	test := new(Arraytest)
	defer cancel()

	if err := c.BindJSON(&test); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}
	newTest := Arraytest{
		ID:         primitive.NewObjectID(),
		First_name: test.First_name,
		Last_name:  test.Last_name,
		Username:   test.Username,
		Email:      test.Email,
		Password:   test.Password,
	}
	result, err := userCollection.InsertOne(ctx, newTest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "success", "Data": map[string]interface{}{"data": result}})

	// c.IndentedJSON(http.StatusOK, gin.H{"message": result})

}

// func postData(c *gin.Context) {
// 	var newData Arraytest

// 	if err := c.BindJSON(&newData); err != nil {
// 		return
// 	}

// 	emp1 = append(emp1, newData)
// 	c.IndentedJSON(http.StatusCreated, newData)

// }

// func getDatabyID(c *gin.Context) {
// 	name := c.Param("name")

// 	for i := 0; i < len(emp1); i++ {
// 		if emp1[i].Name == name {
// 			c.IndentedJSON(http.StatusOK, emp1[i])
// 			return

// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
// }

func main() {

	router := gin.Default()
	connectDB()
	router.POST("/", postData)
	// router.POST("/postData", postData)
	// router.GET("/getDatabyId/:id", getDatabyID)
	router.Run("localhost: 5000")

}

// func main() {
// 	fmt.Println("Hello world")
// 	const Name = "Idowu"
// 	var file = "Paul"
// 	// myArray = [1, 2, 4]
// 	fmt.Println(Name + " " + file)
// 	fmt.Println("What is your name:")

// 	var name string
// 	fmt.Scanln(&name)
// 	fmt.Println("Welcome,", name)
// 	var A = false
// 	println(A)

// type Arraytest struct {
// 	ID     int
// 	Name   string
// 	Salary float32
// }

// var emp1 = []Arraytest{
// 	{ID: 1, Name: "Idowu", Salary: 12345.09},
// 	{ID: 2, Name: "Paul", Salary: 300025.09},
// }

// var p = []string{}

// for i := range emp1 {
// 	p := append(p, emp1[i].Name)
// 	fmt.Print(p)

// 	// fmt.Println(emp1[i].Name)
// }

// for i := 0; i < len(emp1); i++ {
// 	p := append(p, emp1[i].Name)
// 	fmt.Println(p)
// 	fmt.Println(emp1[i].Name)
// }

// fmt.Println(emp1[0].ID)

// // fmt.Println(&Arraytest)

// var myArray = []int{1, 3, 5, 6}
// // myArray2 := []string{"Idowu", "Paul"}

// fmt.Println(myArray)

// aMap := map[string]string{
// 	"Name":   "Idowu",
// 	"Status": "Single",
// 	"Lines":  "Jog",
// }
// fmt.Println(aMap["Name"])

// // for i := 0; i < len(aMap); i++ {
// // 	fmt.Println(aMap[i])

// // }

// for i := range aMap {
// 	fmt.Println(aMap[i])
// }

// for i := 0; i < len(aMap); i++ {
// 	fmt.Println(i)
// }

// for i := range Arraytest{

// }

// fmt.Println(myArray2)
// fmt.Println(myArray)

// for i := 0; i < myArray.count; i++ {

// }
// fmt.Println(myArray)
// }