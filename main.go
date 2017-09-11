package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hashicorp/consul/api"
)

func exitOnErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	var (
		withNodes    = flag.Bool("nodes", true, "Include node names in output")
		withServices = flag.Bool("services", true, "Include services in output")
	)
	flag.Parse()

	if !(*withNodes || *withServices) {
		flag.Usage()
		os.Exit(2)
	}

	client, err := api.NewClient(api.DefaultConfig())
	exitOnErr(err)

	catalog := client.Catalog()

	if *withNodes {
		nodes, _, err := catalog.Nodes(nil)
		exitOnErr(err)

		for _, node := range nodes {
			fmt.Println(node.Node)
		}
	}

	if *withServices {
		services, _, err := catalog.Services(nil)
		exitOnErr(err)

		for service, tags := range services {
			fmt.Println(service)
			for _, tag := range tags {
				fmt.Printf("%s.%s\n", tag, service)
			}
		}
	}
}
