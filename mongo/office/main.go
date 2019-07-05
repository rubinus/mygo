package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
@Time : 2019-07-05 17:58
@Author : rubinus.chu
@File : main.go
@project: mygo
*/

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opts := &options.ClientOptions{}
	opts.SetMaxPoolSize(100)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://182.92.53.129:27017"), opts)
	if err != nil {
		fmt.Println("connection failed", err)
		return
	}
	collection := client.Database("profiles").Collection("bj")

	//find(ctx, collection)
	wg := sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go findOne(ctx, collection, wg)
	}
	wg.Wait()

}

type Res struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Username string             `json:"username"`
}

func findOne(ctx context.Context, collection *mongo.Collection, wg sync.WaitGroup) {
	defer wg.Done()
	result := Res{}
	filter := bson.M{"username": "zhangsan"}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("没有结果")
			return
		}
		fmt.Println(err)
	}

	fmt.Println("成功:", result.Id.Hex(), result.Username)
}

func find(ctx context.Context, collection *mongo.Collection) {
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal("find....", err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result....
		fmt.Println(result, "=====result=====")
	}
	if err := cur.Err(); err != nil {
		log.Fatal("cur.err", err)
	}
}
