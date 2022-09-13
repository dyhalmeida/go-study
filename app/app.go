package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Main: Returns the application of command line ready to be executed
func Main() *cli.App {
	app := cli.NewApp()
	app.Name = "GO IP"
	app.Usage = "Search IP and server names on the web"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Search the IP from a web address",
			Flags: flags,
			Action: findIp,
		},
		{
			Name: "server",
			Usage: "Search the server name on the web",
			Flags: flags,
			Action: findName,
		},
	}

	return app
}

func findIp(c *cli.Context) {
	host := c.String("host")
	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func findName(c *cli.Context) {
	host := c.String("host")
	servers, error := net.LookupNS(host)
	if error != nil {
		log.Fatal(error)
	}
	for _, server := range servers {
		fmt.Println(server.Host)
	}
}