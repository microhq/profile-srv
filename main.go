package main

import (
	"log"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/microhq/profile-srv/db"
	"github.com/microhq/profile-srv/handler"
	"github.com/microhq/profile-srv/proto/record"
)

func main() {

	service := micro.NewService(
		micro.Name("go.micro.srv.profile"),
		micro.Flags(
			cli.StringFlag{
				Name:   "database_url",
				EnvVar: "DATABASE_URL",
				Usage:  "The database URL e.g root@tcp(127.0.0.1:3306)/profile",
			},
		),
		micro.Action(func(c *cli.Context) {
			if len(c.String("database_url")) > 0 {
				db.Url = c.String("database_url")
			}
		}),
	)

	service.Init()

	db.Init()

	record.RegisterRecordHandler(service.Server(), new(handler.Record))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
