package bord

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

func init() {
	SetDefaultWriter(ioutil.Discard)
}

func TestLoggingOn(t *testing.T) {

	bitmask := TurnOnAllLogging()
	if bitmask != 63 {
		t.Fatalf("Bitmask did not set correctly Expected '63' got '%d'\n", bitmask)
	}

	bitmask = SetLogBitmask(63) // 1+2+4+8+16+32=63
	if bitmask != 63 {
		t.Fatalf("Bitmask did not set correctly Expected '63' got '%d'\n", bitmask)
	}

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

	bitmask := SetLogBitmask(0)
	if bitmask != 0 {
		t.Fatalf("Bitmask did not set correctly Expected '0' got '%d'\n", bitmask)
	}

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

// Code stolen from https://talks.golang.org/2014/testing.slide#23
// Subprocess test
// Won't generate test converage for 'Fatal'
func TestFatal(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		Fatal("This is a fatal log message that will exit the program")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestFatal")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()

	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("Fatal exited with %v, want exit status 1", err)
}

func TestSetLogFatal(t *testing.T) {
	bitmask := SetLogFatal(false)

	if bitmask&logFatal != 0 {
		t.Fatalf("Expected fatal logging to be turned off")
	}

	bitmask = SetLogFatal(true)
	if bitmask&logFatal == 0 {
		t.Fatalf("Expected fatal logging to be turned on")
	}
}

func TestUpdateTimeFormat(t *testing.T) {

	defaultFormat := "2006-01-02 15:04:05"
	newFormat := "2006-01-02"

	SetTimeFormat(newFormat)

	if timeFormat != newFormat {
		t.Fatalf("Expected %s got %s", newFormat, timeFormat)
	}

	SetTimeFormat(defaultFormat)

	if timeFormat != defaultFormat {
		t.Fatalf("Expected %s got %s", defaultFormat, timeFormat)
	}

}
