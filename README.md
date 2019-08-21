# Golang MongoDB

## Install and SETUP GO
Download and configure your workspace with latest version of Go and correct environment path.
- [Last Version](https://golang.org/dl/)
- [Windows](http://www.wadewegner.com/2014/12/easy-go-programming-setup-for-windows/)
- [Linux](http://www.tecmint.com/install-go-in-linux/)

## Install Driver
Install dep [klik here](https://golang.github.io/dep/docs/installation.html) and Create new dep
```
dep init 
```
Install driver to build rest and elastic
```
dep ensure -add github.com/gin-gonic/gin
```
```
dep ensure -add go.mongodb.org/mongo-driver
```
## Import
```
import (
	"context"
	"go-mongodb/config"
	"log"
	"net/http"

	"go-mongodb/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)
```
## Struct
```
type Doc struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
```

## Connection
```
func Connect() (*mongo.Collection, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// set collection name
	collection := client.Database("cobaDB").Collection("mytable2")

	return collection, nil
}
```

## SearchAll data from mongoDB
```
func SearchAll(c *gin.Context) {
	var (
		docs []*models.Doc
		r    models.Response
	)

	db, err := config.Connect()
	//set limit data
	findOptions := options.Find()
	findOptions.SetLimit(5)

	searchAll, err := db.Find(context.Background(), bson.D{{}}, findOptions)
	if err != nil {
		r.Status = http.StatusInternalServerError
		r.Message = err.Error()
	}

	for searchAll.Next(context.Background()) {
		var doc models.Doc
		err := searchAll.Decode(&doc)
		if err != nil {
			log.Fatal(err)
		}
		docs = append(docs, &doc)
	}
	if err := searchAll.Err(); err != nil {
		log.Fatal(err)
	}

	r.Status = http.StatusOK
	r.Message = "Success"
	r.Data = docs
	log.Println("Data database")
	c.JSON(http.StatusOK, gin.H{
		"status":  r.Status,
		"message": r.Message,
		"data":    r.Data,
	})
}
```
## Run
```
go run main.go
```
