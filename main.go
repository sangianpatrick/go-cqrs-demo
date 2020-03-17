package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sangianpatrick/go-cqrs-demo/helper/mongodb/repository"
	"go.mongodb.org/mongo-driver/bson"
)

// User struct
type User struct {
	Name string `bson:"name" json:"name"`
}

func main() {
	var user *User
	mongodbClient, err := repository.NewClient("mongodb://localhost:27017/go-cqrs-demo")
	if err != nil {
		log.Fatal(err)
	}

	err = mongodbClient.Connect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	DB := mongodbClient.Database("go-cqrs-demo")
	C := DB.Collection("haha")
	SR := C.FindOne(context.Background(), bson.M{})
	SR.Decode(&user)

	fmt.Println(user)

}
