package users

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/KenethSandoval/fofewi/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	dbc := db.MongoCN.Database("fofewi")
	col := dbc.Collection("users")

	var resultado []*Users
	ID := "61e5cfce77c3d237b2ded8ca"

	condicion := bson.M{
		"_id": ID,
	}

	opciones := options.Find()

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
	}

	for cursor.Next(context.TODO()) {
		var registro Users
		err := cursor.Decode(&registro)
		fmt.Println(&registro)
		if err != nil {
			log.Fatal(err.Error())
		}
		resultado = append(resultado, &registro)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(resultado)
}
