package app

import "github.com/urfave/cli"

// Returns IP and Server Name of some sites
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "App Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidores"
	return app
}
