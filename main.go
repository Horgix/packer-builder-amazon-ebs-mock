package main

import (
	"log"

	"github.com/Horgix/packer-builder-amazon-ebs-mock/amazon-ebs-mock"

	"github.com/hashicorp/packer/packer/plugin"
)

func main() {
	log.Println("Starting...")
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}

	server.RegisterBuilder(amazonebsmock.NewBuilder())
	server.Serve()
}
