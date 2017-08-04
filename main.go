package main

import (
	"log"

	"github.com/horgix/packer-builder-amazon-ebs-mock/amazon-ebs-mock"

	"github.com/hashicorp/packer/packer/plugin"
)

func main() {
	log.Println("Starting...")
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}

	server.RegisterBuilder(new(amazonebsmock.Builder))
	server.Serve()
}
