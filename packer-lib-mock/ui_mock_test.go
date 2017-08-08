package packermock_test

import (
	"fmt"
	"github.com/Horgix/packer-builder-amazon-ebs-mock/packer-lib-mock"
	//"log"
	"testing"
)

func TestSay(t *testing.T) {
	ui := new(packermock.MockUi)

	// Check that it is well initalized at 0
	if ui.SayCount != 0 {
		t.Errorf("Number of calls to ui.Say() was incorrect, "+
			"got %d but should be %d at initialization", ui.SayCount,
			0)
	}

	// We'll call Say() 4 times. TODO: generate some tests for that
	const expectedSayCount = 4
	for i := 1; i <= expectedSayCount; i++ {
		ui.Say(fmt.Sprintf("Call %d", i))
	}

	// Check that every call was rightly recorded
	if ui.SayCount != expectedSayCount {
		t.Errorf("Number of calls to ui.Say() was incorrect, "+
			"got %d but expected %d", ui.SayCount,
			expectedSayCount)
	}
}
