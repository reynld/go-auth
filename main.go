package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/reynld/go-auth/pkg/db"
	"github.com/reynld/go-auth/pkg/server"
)

// getAllEnv makes sure all enviroment variables are set before running
func getAllEnv() error {
	keys := []string{
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_NAME",
		"DB_PASSWORD",
		"ENVIROMENT",
		"PORT",
		"JWT_KEY",
	}

	values := map[string]string{}

	for _, key := range keys {
		v := os.Getenv(key)
		if v == "" {
			return fmt.Errorf("eviroment variable %s is required", key)
		}
		values[key] = v
	}

	return nil
}

func main() {
	godotenv.Load()

	err := getAllEnv()
	if err != nil {
		log.Fatal(err)
	}

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
			db.RunMigrations(s.DB)
		}
		if *seed {
			db.RunSeeds(s.DB)
		}

	} else {
		fmt.Println("pass at least one argument")
		flag.Usage()
	}
}
