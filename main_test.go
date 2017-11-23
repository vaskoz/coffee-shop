package main

import (
	"bytes"
	"os"
	"syscall"
	"testing"
)

func TestBadEnv(t *testing.T) {
	buff := new(bytes.Buffer)
	stderr = buff
	exit = func(code int) {}
	main()
	os.Setenv("COFFEE_SHOP_CLOSE_TIME", "1")
	main()
	os.Setenv("COFFEE_SHOP_SHUTDOWN", "1")
	main()
	os.Setenv("COFFEE_SHOP_CUSTOMERS", "1")
	main()
}

func TestStore(t *testing.T) {
	buff := new(bytes.Buffer)
	stderr = buff
	exit = func(code int) {}
	os.Setenv("COFFEE_SHOP_CLOSE_TIME", "1")
	os.Setenv("COFFEE_SHOP_SHUTDOWN", "0")
	os.Setenv("COFFEE_SHOP_CUSTOMERS", "1")
	os.Setenv("COFFEE_SHOP_BARISTAS", "1")
	main()
}

func TestSignalStore(t *testing.T) {
	buff := new(bytes.Buffer)
	stderr = buff
	exit = func(code int) {}
	stop = make(chan os.Signal, 1)
	os.Setenv("COFFEE_SHOP_CLOSE_TIME", "1")
	os.Setenv("COFFEE_SHOP_SHUTDOWN", "1")
	os.Setenv("COFFEE_SHOP_CUSTOMERS", "1")
	os.Setenv("COFFEE_SHOP_BARISTAS", "1")
	stop <- syscall.SIGINT
	main()
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
