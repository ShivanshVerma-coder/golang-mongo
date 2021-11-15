package database

import (
	"context"
	"fmt"
	"log"

	model "github.com/ShivanshVerma-coder/golang-mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//insert 1 record
func InsertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inserted)

	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

//update 1 record
func UpdateOneMovie(movieId string) {
	fmt.Println(movieId)
	id, _ := primitive.ObjectIDFromHex(movieId)
	fmt.Println("id", id)
	filter := bson.M{"_id": id}
	fmt.Println(filter)
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated 1 movie in db with id: ", result.ModifiedCount)
}

//delete 1 record
func DeleteOneMovie(movieId string) *mongo.DeleteResult {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted 1 movie in db with id: ", deleteCount)
	return deleteCount
}

//delete all movies
func DeleteAllMovies(movieId string) int64 {
	filter := bson.D{{}}

	deleteResult, err := collection.DeleteMany(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted", deleteResult.DeletedCount, "movies")
	return deleteResult.DeletedCount
}

//get all movies
func GetAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		cursor.Decode(&movie)
		movies = append(movies, movie)
	}
	fmt.Println("Movies: ", movies)
	return movies
}
