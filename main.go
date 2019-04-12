package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/reynld/go-auth/pkg/server"
)

func main() {
	godotenv.Load()
	serve := flag.Bool("serve", false, "runs server")
	migrate := flag.Bool("migrate", false, "migrates database")
	seed := flag.Bool("seed", false, "seeds database")
	flag.Parse()

	if len(os.Args) > 1 {

		if flag.NFlag() != 1 {
			fmt.Println("pass just one argument")
			flag.Usage()
			os.Exit(1)
		}

		s := server.Server{}
		s.Initialize()

		if *serve {
			port := os.Getenv("PORT")
			if port == "" {
				log.Fatal("PORT env variable is required\n")
			}
			fmt.Printf("server listening on port: %s\n", port)
			s.Run(fmt.Sprintf(":%s", port))
		}
		if *migrate {
			s.RunMigrations()
		}
		if *seed {
			s.RunSeeds()
		}

	} else {
		fmt.Println("pass at least one argument")
		flag.Usage()
	}
}
