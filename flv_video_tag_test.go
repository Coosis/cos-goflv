package cosgoflv

import (
	"testing"
)

func TestIsValidFrameType(t *testing.T) {
	err := IsValidFrameType(0)
	if err == nil {
		t.Errorf("No error, but 0 is an invalid frame type")
	}
	err = IsValidFrameType(1)
	if err != nil {
		t.Errorf("Error: %v, but 1 is a valid frame type", err)
	}
	err = IsValidFrameType(6)
	if err == nil {
		t.Errorf("No error, but 6 is an invalid frame type")
	}
}

func TestIsValidCodecID(t *testing.T) {
	err := IsValidCodecID(0)
	if err == nil {
		t.Errorf("No error, but 0 is an invalid codec id")
	}
	err = IsValidCodecID(1)
	if err != nil {
		t.Errorf("Error: %v, but 1 is a valid codec id", err)
	}
	err = IsValidCodecID(3)
	if err == nil {
		t.Errorf("No error, but 3 is an invalid codec id")
	}
}
