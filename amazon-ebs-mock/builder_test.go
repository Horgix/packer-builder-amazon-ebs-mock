package amazonebsmock_test

import (
	"fmt"
	"github.com/Horgix/packer-builder-amazon-ebs-mock/amazon-ebs-mock"
	"github.com/Horgix/packer-builder-amazon-ebs-mock/packer-lib-mock"
	"log"
	"testing"
)

func TestLoadDefaultConfig_DefaultGenerate(t *testing.T) {
	const expectedGenerate = true

	builder := new(amazonebsmock.Builder)
	builder.LoadDefaultConfig()

	if builder.Config.Generate != expectedGenerate {
		t.Error(fmt.Sprintf("LoadDefaultConfig() should return a default config with Generate to '%v' but it reported '%v'", expectedGenerate, builder.Config.Generate))
	}
}

//func TestNewBuilder_DefaultAmount(t *testing.T) {
//	const expectedAmount = 1
//
//	builder := new(amazonebsmock.Builder)
//
//        if builder.Config.Amount != expectedAmount {
//          t.Error("NewBuilder() should return a default config with Generate to '%v' but it reported '%v'", expectedAmount, builder.Config.Amount)
//        }
//}

// TODO : Test config parsing

func TestPrepare_Rand(t *testing.T) {
	builder := new(amazonebsmock.Builder)
	builder.Prepare()

	if builder.TestMsg != "Rand seeded" {
		t.Errorf("Prepare() should initialize rand but didn't report it")
	}
}

// func TestPrepare(t *testing.T) {
// 	builder := new(amazonebsmock.Builder)
// 	builder.Prepare()
//
// 	if builder.TestMsg != "Rand seeded" {
// 		t.Errorf("Prepare() should initialize rand but didn't report it")
// 	}
// }

// Dumb test, but you know, coverage.
func TestCancel(t *testing.T) {
	builder := new(amazonebsmock.Builder)
	builder.Cancel()
	if builder.TestMsg != "This method is doing nothing" {
		t.Errorf("Cancel() should be doing nothing but it didn't report it")
	}
}

// Check that Run() method notify the user as expected
func TestRun_UiCalls(t *testing.T) {
	// Initialize and Prepare Builder
	builder := new(amazonebsmock.Builder)
	builder.Prepare()

	// Mock the "ui" part so we can count calls to ui.Say()
	ui := &packermock.MockUi{}

	log.Printf("ui.Say call SayCounter pre Run: %v", ui.SayCount)
	builder.Run(ui, nil, nil)
	log.Printf("ui.Say call SayCounter post Run: %v", ui.SayCount)

	// We should have 4 calls to ui.Say()
	const expectedSayCount = 2
	if ui.SayCount != expectedSayCount {
		t.Errorf("Number of calls to ui.Say() was incorrect, "+
			"got %d but expected %d", ui.SayCount,
			expectedSayCount)
	}
}
