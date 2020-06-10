package main

import (
	"context"
	"fmt"

	"google.golang.org/api/option"

	firebase "firebase.google.com/go"
)

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

}
