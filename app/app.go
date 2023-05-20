package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Returns IP and Server Name of some sites
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "App Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidores"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs de urls na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: searchIps,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
