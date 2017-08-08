package packermock

import (
	"fmt"
	"log"
)

type MockUi struct {
	// Counters for each method call on this struct
	SayCount     int
	AskCount     int
	MessageCount int
	ErrorCount   int
	MachineCount int
}

func (ui *MockUi) Say(message string) {
	log.Printf("ui.Say mock call")
	ui.SayCount++
}

func (ui *MockUi) Ask(query string) (string, error) {
	log.Printf("ui.Ask mock call")
	ui.AskCount++
	return fmt.Sprintf("I am a mocked Response. Your query: %s", query), nil
}

func (ui *MockUi) Message(message string) {
	log.Printf("ui.Message mock call")
	ui.MessageCount++
}

func (ui *MockUi) Error(message string) {
	log.Printf("ui.Error mock call")
	ui.ErrorCount++
}

func (ui *MockUi) Machine(t string, args ...string) {
	log.Printf("ui.Machine mock call")
	ui.MachineCount++
}
