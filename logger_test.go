package bord

import (
	"io/ioutil"
	"testing"
)

func init() {
	SetDefaultWriter(ioutil.Discard)
}

func TestLoggingOn(t *testing.T) {

	SetLogBitmask(31) // 1+2+4+8+16=31
	if ok := Error("Test %s", "error"); !ok {
		t.Fatalf("Logging error failed when it should have worked\n")
	}
	if ok := Warn("Test %s", "warn"); !ok {
		t.Fatalf("Logging warn failed when it should have worked\n")
	}
	if ok := Info("Test %s", "info"); !ok {
		t.Fatalf("Logging info failed when it should have worked\n")
	}
	if ok := Debug("Test %s", "debug"); !ok {
		t.Fatalf("Logging debug failed when it should have worked\n")
	}
	if ok := Custom(ioutil.Discard, "CUSTOM", "Test %s", "custom"); !ok {
		t.Fatalf("Logging custom failed it should have worked\n")
	}
}

func TestLoggingOff(t *testing.T) {

	SetLogBitmask(0)
	if ok := Error("Test %s", "error"); ok {
		t.Fatalf("Logging error worked when it should have failed\n")
	}
	if ok := Warn("Test %s", "warn"); ok {
		t.Fatalf("Logging warn worked when it should have failed\n")
	}
	if ok := Info("Test %s", "info"); ok {
		t.Fatalf("Logging info worked when it should have failed\n")
	}
	if ok := Debug("Test %s", "debug"); ok {
		t.Fatalf("Logging debug worked when it should have failed\n")
	}
	if ok := Custom(ioutil.Discard, "CUSTOM", "Test %s", "custom"); ok {
		t.Fatalf("Logging custom worked it should have failed\n")
	}
}

func TestIndividualOn(t *testing.T) {

	SetLogError(true)
	if ok := Error("Test %s", "error"); !ok {
		t.Fatalf("Logging error failed when it should have worked\n")
	}
	SetLogWarning(true)
	if ok := Warn("Test %s", "warn"); !ok {
		t.Fatalf("Logging warn failed when it should have worked\n")
	}
	SetLogInfo(true)
	if ok := Info("Test %s", "info"); !ok {
		t.Fatalf("Logging info failed when it should have worked\n")
	}
	SetLogDebug(true)
	if ok := Debug("Test %s", "debug"); !ok {
		t.Fatalf("Logging debug failed when it should have worked\n")
	}
	SetLogCustom(true)
	if ok := Custom(ioutil.Discard, "CUSTOM", "Test %s", "custom"); !ok {
		t.Fatalf("Logging custom failed it should have worked\n")
	}
}

func TestIndividualOff(t *testing.T) {

	SetLogError(false)
	if ok := Error("Test %s", "error"); ok {
		t.Fatalf("Logging error worked when it should have failed\n")
	}
	SetLogWarning(false)
	if ok := Warn("Test %s", "warn"); ok {
		t.Fatalf("Logging warn worked when it should have failed\n")
	}
	SetLogInfo(false)
	if ok := Info("Test %s", "info"); ok {
		t.Fatalf("Logging info worked when it should have failed\n")
	}
	SetLogDebug(false)
	if ok := Debug("Test %s", "debug"); ok {
		t.Fatalf("Logging debug worked when it should have failed\n")
	}
	SetLogCustom(false)
	if ok := Custom(ioutil.Discard, "CUSTOM", "Test %s", "custom"); ok {
		t.Fatalf("Logging custom worked it should have failed\n")
	}
}
