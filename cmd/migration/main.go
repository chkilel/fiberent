package main

import (
	"context"

	"log"

	"github.com/chkilel/fiberent/ent"
	"github.com/chkilel/fiberent/ent/migrate"
	"github.com/chkilel/fiberent/infrastructure/ent/datastore"
	"github.com/chkilel/fiberent/pkg/config"
)

func main() {
	config.LoadEnvironmentFile(".env")

	client, err := datastore.NewClient()
	if err != nil {
		log.Fatalf("failed opening Postgres client: %v", err)
	}
	defer client.Close()
	createDBSchema(client)
}

func createDBSchema(client *ent.Client) {
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
