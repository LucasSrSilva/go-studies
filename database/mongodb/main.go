package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	uri := "mongodb://admin:admin@localhost:27017/admin?authSource=admin"

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Print("Erro ao conectar", err)
		os.Exit(1)
	}
	defer client.Disconnect(context.TODO())
	coll := client.Database("netflocos").Collection("filmes")
	novoFilme := bson.D{
		{Key: "title", Value: "Sem volta para o futuro"},
		{Key: "diretor", Value: "Robert Zemeckis"},
		{Key: "ano", Value: 1985},
		{Key: "generos", Value: []string{"Ficção Científica", "Aventura"}},
	}

	// 2. Inserindo no banco de dados
	insertRes, err := coll.InsertOne(context.TODO(), novoFilme)
	if err != nil {
		fmt.Println("Erro ao inserir documento:", err)
		os.Exit(1)
	}

	// 3. O MongoDB gera automaticamente um ID único (_id) para cada documento
	fmt.Printf("Documento inserido com sucesso! ID: %v\n", insertRes.InsertedID)

	title := "Sem volta para o futuro"

	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{Key: "title", Value: title}}).
		Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}

// docker run --name my-mongodb \
//   -p 27017:27017 \
//   -e MONGO_INITDB_ROOT_USERNAME=admin \
//   -e MONGO_INITDB_ROOT_PASSWORD=admin \
//   -v mongo_data:/data/db \
//   -d mongo:latest
