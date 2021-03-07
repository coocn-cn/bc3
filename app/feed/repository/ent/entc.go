// +build ignore

package main

import (
	"context"
	"log"
	"os"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/coocn-cn/bc3/app/feed/repository/ent"
	_ "github.com/lib/pq"
)

func generate() {
	opts := []entc.Option{
		entc.FeatureNames("privacy"),
	}

	err := entc.Generate("./schema", &gen.Config{
		// Templates: entgql.AllTemplates,
	}, opts...)
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

func migrate() {
	client, err := ent.Open("postgres", "postgresql://postgres:postgres@127.0.0.1/postgres")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Dump migration changes to stdout.
	if err := client.Schema.WriteTo(ctx, os.Stdout); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
}

func main() {
	cmd := "generate"
	if len(os.Args) == 2 {
		cmd = os.Args[1]
	}

	switch cmd {
	case "migrate":
		migrate()
	default:
		generate()
	}
}
