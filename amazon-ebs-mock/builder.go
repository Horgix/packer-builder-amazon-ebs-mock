package amazonebsmock

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	awscommon "github.com/hashicorp/packer/builder/amazon/common"
	"github.com/hashicorp/packer/common"
	"github.com/hashicorp/packer/helper/config"
	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/template/interpolate"
)

const BuilderId = "horgix.amazonebsmock"

// Hello Go, we can't even declare const slices...
var regions = []string{
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

// TODO : ideally, the Config and initialization should be the same than the
// official Amazon EBS Builder
type Config struct {
	common.PackerConfig `mapstructure:",squash"`

	Generate bool   `mapstructure:"generate"`
	Amount   int    `mapstructure:"amount"`
	Region   string `mapstructure:"region"`

	ctx interpolate.Context
}

type Builder struct {
	Config Config
	// The testMsg field used to get some feedback during tests from method
	// which don't return anything like Cancel()
	TestMsg string
}

func (b *Builder) LoadDefaultConfig() {
	log.Println("LoadDefaultConfig(): Initializing builder's configuration")
	b.Config = Config{}
	log.Println(b.Config)
	log.Println("LoadDefaultConfig(): Setting default: Generate = false")
	b.Config.Generate = true
	log.Println("LoadDefaultConfig(): Setting default: Amount = 1")
	b.Config.Amount = 1
	log.Println("LoadDefaultConfig(): Setting default: Region = \"\"")
	b.Config.Region = ""
}

func (b *Builder) Prepare(raws ...interface{}) ([]string, error) {
	log.Println("Hello I'm a custom builder and this is the Prepare step")

	b.LoadDefaultConfig()
	log.Println("Prepare(): initial configuration:")
	log.Println("  - Generate: ", b.Config.Generate)
	log.Println("  - Amount: ", b.Config.Amount)
	log.Println("  - Region: \"" + b.Config.Region + "\"")
	log.Println("Prepare(): parsing configuration")
	err := config.Decode(&b.Config, &config.DecodeOpts{Interpolate: false}, raws...)
	if err != nil {
		return nil, err
	}
	log.Println("Prepare(): parsed configuration with success")
	log.Println("Prepare(): current configuration:")
	log.Println("  - Generate: ", b.Config.Generate)
	log.Println("  - Amount: ", b.Config.Amount)
	log.Println("  - Region: \"" + b.Config.Region + "\"")

	log.Println("Prepare(): checking for incorrect parameters")
	if b.Config.Amount != 1 && b.Config.Region != "" {
		// TODO : Improve this error, maybe use Packer's types?
		return nil, errors.New("FAIL")
	}

	log.Println("Prepare(): Initializing random generator...")
	log.Println(b.Config.Generate)
	rand.Seed(time.Now().Unix())
	b.TestMsg = "Rand seeded"

	return nil, nil
}

func (b *Builder) Run(ui packer.Ui, hook packer.Hook, cache packer.Cache) (packer.Artifact, error) {
	log.Println("Run(): Hello I'm a custom builder")

	amis := make(map[string]string)

	ui.Say(fmt.Sprintf("Generating mock Artifact (%d AMI IDs)...", b.Config.Amount))
	for i := 0; i < b.Config.Amount; i++ {
		amiId := fmt.Sprintf("ami-%06d", i+1)
		if b.Config.Generate {
			// Packer is not able to build 2 AMIs in the same
			// region in the same Builder, so:
			// - We take a random region
			// - We remove it from the pool

			// Take a random region
			idx := rand.Intn(len(regions))
			region := regions[idx]

			// Remove it from the pool
			regions[idx] = regions[len(regions)-1]
			regions = regions[:len(regions)-1]

			amis[region] = amiId
		} else {
			amis["eu-west-1"] = amiId
		}
	}

	log.Println("Run(): tranforming into official module artifact type...")
	artifact := &awscommon.Artifact{
		Amis:           amis,
		BuilderIdValue: BuilderId,
	}

	ui.Say("Generated mock Artifact with success :)")
	return artifact, nil
}

func (b *Builder) Cancel() {
	b.TestMsg = "This method is doing nothing"
}
