package main

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/reynld/go-auth/server"
	"github.com/reynld/go-auth/server/models"
)

func main() {
	migrate := flag.Bool("migrate", false, "migrates database")
	seed := flag.Bool("seed", false, "seeds database")
	seedCache := flag.Bool("seed-cache", false, "seeds redis cache")
	run := flag.Bool("run", false, "runs server")
	flag.Parse()

	if len(os.Args) > 1 {
		if flag.NFlag() != 1 {
			fmt.Println("pass just one argument")
			flag.Usage()
			os.Exit(1)
		}

		godotenv.Load()
		s := server.Server{}
		err := s.CheckEnviroment()
		s.Panic(err)

		s.Initialize()

		if *run {
			s.Run()
		}
		if *migrate {
			models.RunMigrations(s.DB)
		}
		if *seed {
			models.RunDBSeeds(s.DB)
		}
		if *seedCache {
			models.RunCacheSeeds(s.Cache)
		}

	} else {
		fmt.Println("pass at least one argument")
		flag.Usage()
		os.Exit(1)
	}

}
