package main

import (
	"context"
	"log"

	"github.com/khozaei/healthio/ent"
	"github.com/khozaei/healthio/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "root:252@/healthiodb?parseTime=true")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	//	defer client.Close()
	// Run the auto migration tool.
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	routes.SetHandlerConfig(routes.HandlerConfig{
		Client:  client,
		Context: ctx,
	})
	routes.Run()
}
