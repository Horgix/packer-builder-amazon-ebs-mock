package amazonebsmock

import (
	"log"
	"math/rand"
	"time"

	awscommon "github.com/hashicorp/packer/builder/amazon/common"
	"github.com/hashicorp/packer/common"
	"github.com/hashicorp/packer/packer"
)

const BuilderId = "horgix.amazonebsmock"

// Hello Go, we can't even declare const slices...
var regions = [...]string{
	"us-east-2",
	"us-east-1",
	"us-west-1",
	"us-west-2",
	"ca-central-1",
	"ap-south-1",
	"ap-northeast-2",
	"ap-southeast-1",
	"ap-southeast-2",
	"ap-northeast-1",
	"eu-central-1",
	"eu-west-1",
	"eu-west-2",
	"sa-east-1",
}

type Config struct {
	common.PackerConfig `mapstructure:",squash"`
}

type Builder struct {
	config *Config
}

func (b *Builder) Prepare(raws ...interface{}) ([]string, error) {
	log.Println("Hello I'm a custom builder and this is the Prepare step")
	log.Println("Initializing random generator...")
	rand.Seed(time.Now().Unix())
	return nil, nil
}

func (b *Builder) Run(ui packer.Ui, hook packer.Hook, cache packer.Cache) (packer.Artifact, error) {
	log.Println("Hello I'm a custom builder")

	ui.Say("I'm doing nothing...")
	ui.Say("Nothing done with success!")

	ui.Say("Generating mock Artifact...")

	amis := make(map[string]string)

	for i := 0; i < 10; i++ {
		amis[regions[rand.Intn(len(regions))]] = "ami-12345678"
	}

	artifact := &awscommon.Artifact{
		Amis: amis,
	}
	ui.Say("Generated mock Artifact with success :)")
	return artifact, nil
}

func (b *Builder) Cancel() {
}
