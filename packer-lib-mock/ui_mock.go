package packermock

import (
	"log"
)

type MockUi struct {
	SayCount int
}

func (ui *MockUi) Say(message string) {
	log.Printf("ui.Say mock call")
	ui.SayCount++
}

func (ui *MockUi) Ask(query string) (string, error) {
	log.Printf("ui.Ask mock call")
	return "", nil
}

func (ui *MockUi) Message(message string) {
	log.Printf("ui.Message mock call")
}

func (ui *MockUi) Error(message string) {
	log.Printf("ui.Error mock call")
}

func (u *MockUi) Machine(t string, args ...string) {
	log.Printf("ui.Machine mock call")
}
