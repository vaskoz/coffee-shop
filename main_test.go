package main

import (
	"bytes"
	"os"
	"strings"
	"syscall"
	"testing"
)

func TestBadEnv(t *testing.T) {
	buff := new(bytes.Buffer)
	stderr = buff
	var code int
	exit = func(c int) {
		code = c
	}
	main()
	if code != 1 {
		t.Error("expected exit status 1")
	}
	if !strings.Contains(buff.String(), "Bad COFFEE_SHOP_CLOSE_TIME") {
		t.Error("Didn't provide close time, but didn't error")
	}
	code = 0
	os.Setenv("COFFEE_SHOP_CLOSE_TIME", "1")
	main()
	if code != 1 {
		t.Error("expected exit status 1")
	}
	if !strings.Contains(buff.String(), "Bad COFFEE_SHOP_SHUTDOWN") {
		t.Error("Didn't provide shutdown time, but didn't error")
	}
	code = 0
	os.Setenv("COFFEE_SHOP_SHUTDOWN", "1")
	main()
	if code != 1 {
		t.Error("expected exit status 1")
	}
	if !strings.Contains(buff.String(), "Bad COFFEE_SHOP_CUSTOMERS") {
		t.Error("Didn't provide number of customers, but didn't error")
	}
	code = 0
	os.Setenv("COFFEE_SHOP_CUSTOMERS", "1")
	main()
	if code != 1 {
		t.Error("expected exit status 1")
	}
	if !strings.Contains(buff.String(), "Bad COFFEE_SHOP_BARISTAS") {
		t.Error("Didn't provide number of baristas, but didn't error")
	}
}

func TestStore(t *testing.T) {
	buff := new(bytes.Buffer)
	stderr = buff
	var code int = -1
	exit = func(c int) {
		code = c
	}
	os.Setenv("COFFEE_SHOP_CLOSE_TIME", "1")
	os.Setenv("COFFEE_SHOP_SHUTDOWN", "0")
	os.Setenv("COFFEE_SHOP_CUSTOMERS", "1")
	os.Setenv("COFFEE_SHOP_BARISTAS", "1")
	main()
	if code != 0 {
		t.Error("Expected 0 exit status")
	}
	if s := buff.String(); !strings.Contains(s, "Store is closing") || !strings.Contains(s, "Customer 1 says Yum and thanks to Barista 1") {
		t.Errorf("Output was not expected got %v", s)
	}
}

func TestSignalStore(t *testing.T) {
	buff := new(bytes.Buffer)
	stderr = buff
	var code int = -1
	exit = func(c int) {
		code = c
	}
	stop = make(chan os.Signal, 1)
	os.Setenv("COFFEE_SHOP_CLOSE_TIME", "1")
	os.Setenv("COFFEE_SHOP_SHUTDOWN", "2")
	os.Setenv("COFFEE_SHOP_CUSTOMERS", "1")
	os.Setenv("COFFEE_SHOP_BARISTAS", "1")
	stop <- syscall.SIGINT
	main()
	if code != 0 {
		t.Error("Expected 0 exit code")
	}
	if s := buff.String(); !strings.Contains(s, "I received a signal to close the store") || !strings.Contains(s, "Store closed") {
		t.Error("Expected a store to close with all customers served")
	}
}

func TestSignalStoreShutdownTimeout(t *testing.T) {
	buff := new(bytes.Buffer)
	stderr = buff
	var code int = -1
	exit = func(c int) {
		code = c
	}
	stop = make(chan os.Signal, 1)
	os.Setenv("COFFEE_SHOP_CLOSE_TIME", "1")
	os.Setenv("COFFEE_SHOP_SHUTDOWN", "1")
	os.Setenv("COFFEE_SHOP_CUSTOMERS", "1")
	os.Setenv("COFFEE_SHOP_BARISTAS", "1")
	stop <- syscall.SIGINT
	main()
	if code != 0 {
		t.Error("Expected 0 exit code")
	}
	if s := buff.String(); !strings.Contains(s, "I received a signal to close the store") || !strings.Contains(s, "Shutdown time reached") {
		t.Error("Expected a store to close with all customers served")
	}
}

func TestCantOpenStore(t *testing.T) {
	buff := new(bytes.Buffer)
	stderr = buff
	exit = func(code int) {}
	os.Setenv("COFFEE_SHOP_CLOSE_TIME", "0")
	os.Setenv("COFFEE_SHOP_SHUTDOWN", "0")
	os.Setenv("COFFEE_SHOP_CUSTOMERS", "0")
	os.Setenv("COFFEE_SHOP_BARISTAS", "0")
	main()
}
