package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"google.golang.org/api/option"

	firebase "firebase.google.com/go"
)

type MagicalGirls []MagicalGirl

func UnmarshalMagicalGirls(data []byte) (MagicalGirls, error) {
	var r MagicalGirls
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MagicalGirls) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MagicalGirl struct {
	Key       string `json:"key"`
	Name      string `json:"name"`
	Attribute string `json:"attribute"`
	Type      string `json:"type"`
	Status    Status `json:"status"`
}

type Status struct {
	HP      int64 `json:"hp"`
	Attack  int64 `json:"attack"`
	Defense int64 `json:"defense"`
}

func main() {
	// Use a service account
	ctx := context.Background()
	sa := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		fmt.Printf("error initializing app: %v", err)
		return
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		fmt.Printf("error create Firestore client: %v", err)
		return
	}
	defer client.Close()

	// load json file
	bytes, err := ioutil.ReadFile("path/to/magicalGirls.json")
	if err != nil {
		fmt.Printf("error load magicalGirls.json : %v", err)
	}
	magicalGirls, err := UnmarshalMagicalGirls(bytes)
	if err != nil {
		fmt.Printf("error unmarshal magicalGirls.json : %v", err)
	}

	for i := range magicalGirls {
		_, err = client.Collection("private/v1/magicalGirls").Doc(magicalGirls[i].Key).Set(ctx, magicalGirls[i])
		if err != nil {
			fmt.Printf("error adding magicalGirl: %v", err)
			return
		}
	}

}
