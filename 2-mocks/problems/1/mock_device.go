package Mock

import . "github.com/arowden/unit_testing/2-mocks/problems/1/device"

// Create a mock for the Device interface (defined in the device directory).
//
// Running 'go build' in this directory will not result in any error if the Device
// interface has been implemented. You should also record which methods were invoked
// and have a way to assign return values.

// Will cause compile failure if the Mock struct does not implement the Device interface.
var _ Device = &Mock{}

type Mock struct{}
