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

func TestIndividual(t *testing.T) {

	LogError(true)
	if ok := Error("Test %s", "error"); !ok {
		t.Fatalf("Logging error failed when it should have worked\n")
	}
	LogWarning(true)
	if ok := Warn("Test %s", "warn"); !ok {
		t.Fatalf("Logging warn failed when it should have worked\n")
	}
	LogInfo(true)
	if ok := Info("Test %s", "info"); !ok {
		t.Fatalf("Logging info failed when it should have worked\n")
	}
	LogDebug(true)
	if ok := Debug("Test %s", "debug"); !ok {
		t.Fatalf("Logging debug failed when it should have worked\n")
	}
	LogCustom(true)
	if ok := Custom(ioutil.Discard, "CUSTOM", "Test %s", "custom"); !ok {
		t.Fatalf("Logging custom failed it should have worked\n")
	}

	LogError(false)
	if ok := Error("Test %s", "error"); ok {
		t.Fatalf("Logging error worked when it should have failed\n")
	}
	LogWarning(false)
	if ok := Warn("Test %s", "warn"); ok {
		t.Fatalf("Logging warn worked when it should have failed\n")
	}
	LogInfo(false)
	if ok := Info("Test %s", "info"); ok {
		t.Fatalf("Logging info worked when it should have failed\n")
	}
	LogDebug(false)
	if ok := Debug("Test %s", "debug"); ok {
		t.Fatalf("Logging debug worked when it should have failed\n")
	}
	LogCustom(false)
	if ok := Custom(ioutil.Discard, "CUSTOM", "Test %s", "custom"); ok {
		t.Fatalf("Logging custom worked it should have failed\n")
	}
}
