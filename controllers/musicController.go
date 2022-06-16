package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/duduBTW/karaokev2backend/db"
	middlewares "github.com/duduBTW/karaokev2backend/handlers"
	"github.com/duduBTW/karaokev2backend/models"
	"github.com/duduBTW/karaokev2backend/validators"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var client = db.Dbconnect()

const dbName = "Karaokev2"
const colName = "Musics"

var CreateMusicEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var music models.Music
	err := json.NewDecoder(request.Body).Decode(&music)
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}

	if ok, erros := validators.ValidateInputs(music); !ok {
		middlewares.ValidationResponse(erros, response)
		return
	}

	collection := client.Database(dbName).Collection(colName)
	result, err := collection.InsertOne(context.TODO(), music)
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}

	res, _ := json.Marshal(result.InsertedID)
	middlewares.SuccessResponse(`Inserted at `+strings.Replace(string(res), `"`, ``, 2), response)
})

var GetMusicsSectionEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	collection := client.Database(dbName).Collection(colName)

	// Find data
	cursor, err := collection.Find(context.TODO(), bson.M{"section": params["section_id"]})
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}

	defer cursor.Close(context.TODO())
	var musics []primitive.M

	for cursor.Next(context.TODO()) {
		var music bson.M
		err := cursor.Decode(&music)
		if err != nil {
			middlewares.ServerErrResponse(err.Error(), response)
			return
		}

		musics = append(musics, music)
	}

	middlewares.SuccessRespond(musics, response)
})
