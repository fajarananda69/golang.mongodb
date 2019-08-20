package controllers

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

func InsertDB(c *gin.Context) {

	var docs models.Doc
	var result models.Response

	if err := c.BindJSON(&docs); err != nil {
		panic("Malformed request body")
		return
	}

	db, err := config.Connect()
	insertResult, err := db.InsertOne(context.Background(), &docs)
	if err != nil {
		result.Status = http.StatusInternalServerError
		result.Message = err.Error()
	}
	result.Status = http.StatusOK
	result.Message = "Insert Success "
	result.Data = docs
	log.Println("Insert data to database", insertResult.InsertedID)
	c.JSON(http.StatusOK, gin.H{
		"status":  result.Status,
		"message": result.Message,
		"data":    result.Data,
	})

}

func SearchById(c *gin.Context) {

	ID := c.Param("id")
	var (
		r    models.Response
		docs []*models.Doc
	)

	db, err := config.Connect()
	//set limit data
	findOptions := options.Find()
	findOptions.SetLimit(2)
	searchBD, err := db.Find(context.Background(), bson.M{"id": ID}, findOptions)
	if err != nil {
		r.Status = http.StatusInternalServerError
		r.Message = err.Error()
	}

	for searchBD.Next(context.Background()) {
		var doc models.Doc
		err := searchBD.Decode(&doc)
		if err != nil {
			log.Fatal(err)
		}
		docs = append(docs, &doc)
	}
	if err := searchBD.Err(); err != nil {
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

func UpdateDB(c *gin.Context) {
	ID := c.Query("id")

	var (
		doc models.Doc
		r   models.Response
	)

	db, err := config.Connect()

	if err := c.BindJSON(&doc); err != nil {
		log.Fatal(err)
		return
	}

	updateDB, err := db.UpdateOne(context.Background(), bson.M{"id": ID}, bson.M{"$set": &doc})
	if err != nil {
		log.Fatal(err)
	}

	r.Status = http.StatusOK
	r.Message = "Success"
	r.Data = doc
	log.Println("Update data", updateDB.MatchedCount)
	c.JSON(http.StatusOK, gin.H{
		"status":   r.Status,
		"messagae": r.Message,
		"data":     r.Data,
	})
}

func DeleteDB(c *gin.Context) {

	ID := c.Query("id")

	db, err := config.Connect()

	deleteDB, err := db.DeleteOne(context.Background(), bson.M{"id": ID})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Delete data", deleteDB.DeletedCount)
}
