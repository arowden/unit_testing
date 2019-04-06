package mock

import . "github.com/arowden/unit_testing/2-mocks/problems/1/device"

// You've been given a struct Mock that is empty. Make it implement the Device interface
// (defined in the device directory).
//
// If done correctly, running 'go build' in this directory will not result in any error.

// Will cause compile failure if the Mock struct does not implement the Device interface.
var _ Device = &Mock{}

type Mock struct {
	ExecuteCallReceivesCmd   string
	ExecuteCallReturnsOutput string
	ExecuteCallReturnsErr    error
	CloseCallInvoked         bool
	CloseCallReturnsErr      error
}

func (m *Mock) Execute(cmd string) (string, error) {
	m.ExecuteCallReceivesCmd = cmd
	return m.ExecuteCallReturnsOutput, m.ExecuteCallReturnsErr
}

func (m *Mock) Close() error {
	m.CloseCallInvoked = true
	return m.CloseCallReturnsErr
}
