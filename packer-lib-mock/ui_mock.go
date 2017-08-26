package packermock

import (
	"fmt"
	"log"
)

// MockUI is a mock implementation of Packer's Ui interface (packer/ui.go)
// It only includes counter to each of its methods, which will not do anything
type MockUI struct {
	SayCount     int
	AskCount     int
	MessageCount int
	ErrorCount   int
	MachineCount int
}

// Say counts calls to itself. Mock of the Packer Ui.Say() method
func (ui *MockUI) Say(message string) {
	log.Printf("ui.Say mock call")
	ui.SayCount++
}

// Ask counts calls to itself. Mock of the Packer Ui.Ask() method
func (ui *MockUI) Ask(query string) (string, error) {
	log.Printf("ui.Ask mock call")
	ui.AskCount++
	return fmt.Sprintf("I am a mocked Response. Your query: %s", query), nil
}

// Message counts calls to itself. Mock of the Packer Ui.Message() method
func (ui *MockUI) Message(message string) {
	log.Printf("ui.Message mock call")
	ui.MessageCount++
}

// Error counts calls to itself. Mock of the Packer Ui.Error() method
func (ui *MockUI) Error(message string) {
	log.Printf("ui.Error mock call")
	ui.ErrorCount++
}

// Machine counts calls to itself. Mock of the Packer Ui.Machine() method
func (ui *MockUI) Machine(t string, args ...string) {
	log.Printf("ui.Machine mock call")
	ui.MachineCount++
}
