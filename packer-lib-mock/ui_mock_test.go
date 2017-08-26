package packermock_test

import (
	"fmt"
	"github.com/Horgix/packer-builder-amazon-ebs-mock/packer-lib-mock"
	"log"
	"testing"
)

func TestSay(t *testing.T) {
	ui := new(packermock.MockUI)

	// Check that it is well initialized at 0
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

func TestAsk(t *testing.T) {
	ui := new(packermock.MockUI)

	// Check that it is well initialized at 0
	if ui.AskCount != 0 {
		t.Errorf("Number of calls to ui.Ask() was incorrect, "+
			"got %d but should be %d at initialization", ui.AskCount,
			0)
	}

	// We'll call Ask() 4 times. TODO: generate some tests for that
	const expectedAskCount = 4
	for i := 1; i <= expectedAskCount; i++ {
		query := fmt.Sprintf("Call %d", i)
		response, _ := ui.Ask(query)
		expectedResponse := fmt.Sprintf("I am a mocked Response. Your query: %s", query)
		if response != expectedResponse {
			log.Printf("Expected: ", expectedResponse)
			log.Printf("Got: ", response)
			t.Errorf("Didnt get the expected Response")
		}
	}

	// Check that every call was rightly recorded
	if ui.AskCount != expectedAskCount {
		t.Errorf("Number of calls to ui.Ask() was incorrect, "+
			"got %d but expected %d", ui.AskCount,
			expectedAskCount)
	}
}

func TestMessage(t *testing.T) {
	ui := new(packermock.MockUI)

	// Check that it is well initialized at 0
	if ui.MessageCount != 0 {
		t.Errorf("Number of calls to ui.Message() was incorrect, "+
			"got %d but should be %d at initialization", ui.MessageCount,
			0)
	}

	// We'll call Message() 4 times. TODO: generate some tests for that
	const expectedMessageCount = 4
	for i := 1; i <= expectedMessageCount; i++ {
		ui.Message(fmt.Sprintf("Call %d", i))
	}

	// Check that every call was rightly recorded
	if ui.MessageCount != expectedMessageCount {
		t.Errorf("Number of calls to ui.Message() was incorrect, "+
			"got %d but expected %d", ui.MessageCount,
			expectedMessageCount)
	}
}

func TestError(t *testing.T) {
	ui := new(packermock.MockUI)

	// Check that it is well initialized at 0
	if ui.ErrorCount != 0 {
		t.Errorf("Number of calls to ui.Error() was incorrect, "+
			"got %d but should be %d at initialization", ui.ErrorCount,
			0)
	}

	// We'll call Error() 4 times. TODO: generate some tests for that
	const expectedErrorCount = 4
	for i := 1; i <= expectedErrorCount; i++ {
		ui.Error(fmt.Sprintf("Call %d", i))
	}

	// Check that every call was rightly recorded
	if ui.ErrorCount != expectedErrorCount {
		t.Errorf("Number of calls to ui.Error() was incorrect, "+
			"got %d but expected %d", ui.ErrorCount,
			expectedErrorCount)
	}
}

func TestMachine(t *testing.T) {
	ui := new(packermock.MockUI)

	// Check that it is well initialized at 0
	if ui.MachineCount != 0 {
		t.Errorf("Number of calls to ui.Machine() was incorrect, "+
			"got %d but should be %d at initialization", ui.MachineCount,
			0)
	}

	// We'll call Machine() 4 times. TODO: generate some tests for that
	const expectedMachineCount = 4
	for i := 1; i <= expectedMachineCount; i++ {
		ui.Machine(fmt.Sprintf("Call %d", i))
	}

	// Check that every call was rightly recorded
	if ui.MachineCount != expectedMachineCount {
		t.Errorf("Number of calls to ui.Machine() was incorrect, "+
			"got %d but expected %d", ui.MachineCount,
			expectedMachineCount)
	}
}
