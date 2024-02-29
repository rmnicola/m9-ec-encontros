---
title: MongoDB Atlas com Go
sidebar_position: 3
sidebar_class_name: opcional
slug: /altas-go
---

import Admonition from '@theme/Admonition';

# Utilizando o MongoDB Atlas c/ Go

Para interagir com um banco de dados MongoDB Atlas usando Go, você pode usar o
pacote `mongo-go-driver`. Abaixo está um exemplo simples que mostra como se
conectar a um cluster do MongoDB Atlas, inserir um documento em uma coleção e
recuperar documentos da coleção.

Primeiro, certifique-se de ter o `mongo-go-driver` instalado. Se não tiver,
você pode instalá-lo usando o seguinte comando:

```bash
go get go.mongodb.org/mongo-driver/mongo
```

Em seguida, use o seguinte código Go como exemplo:

```go
package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Substitua 'your_connection_string' pela sua string de conexão do MongoDB Atlas
	connectionString := "your_connection_string"

	// Conectar ao cluster do MongoDB Atlas
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

	// Selecionar o banco de dados
	db := client.Database("nome_do_banco_de_dados")

	// Selecionar a coleção
	collection := db.Collection("nome_da_coleção")

	// Inserir um documento na coleção
	document := bson.D{{"nome", "João"}, {"idade", 30}}
	_, err = collection.InsertOne(context.TODO(), document)
	if err != nil {
		log.Fatal(err)
	}

	// Recuperar documentos da coleção
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
}
```

Certifique-se de substituir `'your_connection_string'` pela sua string de
conexão do MongoDB Atlas, que você pode obter do painel do Atlas. Além disso,
substitua `'nome_do_banco_de_dados'` e `'nome_da_coleção'` pelos nomes
apropriados do seu banco de dados e coleção, respectivamente.
