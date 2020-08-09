package main

import (
	"log"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/cipta-ageung/auth-srv/db"
	"github.com/cipta-ageung/auth-srv/db/mysql"
	"github.com/cipta-ageung/auth-srv/handler"
	account "github.com/cipta-ageung/auth-srv/proto/account"
	oauth2 "github.com/cipta-ageung/auth-srv/proto/oauth2"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.auth"),
		micro.Flags(
			cli.StringFlag{
				Name:   "database_url",
				EnvVar: "DATABASE_URL",
				Usage:  "The database URL e.g root@tcp(127.0.0.1:3306)/auth",
			},
		),

		micro.Action(func(c *cli.Context) {
			if len(c.String("database_url")) > 0 {
				mysql.Url = c.String("database_url")
			}
		}),
	)

	// initialise service
	service.Init()

	// register account handler
	account.RegisterAccountHandler(service.Server(), new(handler.Account))
	// register oauth2 handler
	oauth2.RegisterOauth2Handler(service.Server(), new(handler.Oauth2))

	// initialise database
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}