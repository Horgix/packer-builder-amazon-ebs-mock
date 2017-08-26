package packermock

import (
	"fmt"
	"log"
)

type MockUI struct {
	// Counters for each method call on this struct

	SayCount     int
	AskCount     int
	MessageCount int
	ErrorCount   int
	MachineCount int
}

func (ui *MockUI) Say(message string) {
	log.Printf("ui.Say mock call")
	ui.SayCount++
}

func (ui *MockUI) Ask(query string) (string, error) {
	log.Printf("ui.Ask mock call")
	ui.AskCount++
	return fmt.Sprintf("I am a mocked Response. Your query: %s", query), nil
}

func (ui *MockUI) Message(message string) {
	log.Printf("ui.Message mock call")
	ui.MessageCount++
}

func (ui *MockUI) Error(message string) {
	log.Printf("ui.Error mock call")
	ui.ErrorCount++
}

func (ui *MockUI) Machine(t string, args ...string) {
	log.Printf("ui.Machine mock call")
	ui.MachineCount++
}
