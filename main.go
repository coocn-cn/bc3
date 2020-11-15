package main

import (
	"fmt"
	"log"

	"github.com/coocn-cn/bc3/app/task/api/echo"
	"github.com/coocn-cn/bc3/app/task/ent"
	_ "github.com/lib/pq"
)

func main() {
	opts := []ent.Option{
		ent.Debug(),
	}

	client, err := ent.Open("postgres", "host=39.96.11.125 port=5432 sslmode=disable user=postgres password=postgresql+20201111 dbname=basecamp", opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	fmt.Println("http://localhost:8888")
	echo.NewServer(client).Start(":8888")
}
