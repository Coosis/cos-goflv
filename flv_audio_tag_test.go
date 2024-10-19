package cosgoflv

import (
	"testing"
)

func TestIsValidSoundFormat(t *testing.T) {
	err := IsValidSoundFormat(7)
	if err == nil {
		t.Errorf("7 is reserved value, got nothing")
	}
	err = IsValidSoundFormat(8)
	if err == nil {
		t.Errorf("8 is reserved value, got nothing")
	}
	err = IsValidSoundFormat(14)
	if err == nil {
		t.Errorf("14 is reserved value, got nothing")
	}
	err = IsValidSoundFormat(15)
	if err == nil {
		t.Errorf("15 is reserved value, got nothing")
	}
	err = IsValidSoundFormat(16)
	if err == nil {
		t.Errorf("16 is invalid sound format, got nothing")
	}
	err = IsValidSoundFormat(4)
	if err != nil {
		t.Errorf("4 is valid sound format, got error: %v", err)
	}
}

func TestIsValidSamplingRate(t *testing.T) {
	err := IsValidSamplingRate(4)
	if err == nil {
		t.Errorf("4 is invalid sampling rate, got nothing")
	}

	err = IsValidSamplingRate(1)
	if err != nil {
		t.Errorf("1 is valid sampling rate, got error: %v", err)
	}
}

func TestIsValidSoundSize(t *testing.T) {
	err := IsValidSoundSize(2)
	if err == nil {
		t.Errorf("2 is invalid sound size, got nothing")
	}

	err = IsValidSoundSize(1)
	if err != nil {
		t.Errorf("1 is valid sound size, got error: %v", err)
	}
}

func TestIsValidSoundType(t *testing.T) {
	err := IsValidSoundType(2)
	if err == nil {
		t.Errorf("2 is invalid sound type, got nothing")
	}

	err = IsValidSoundType(1)
	if err != nil {
		t.Errorf("1 is valid sound type, got error: %v", err)
	}
}
