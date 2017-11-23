package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var (
	stop             = make(chan os.Signal, 1)
	exit             = os.Exit
	stderr io.Writer = os.Stderr
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	logger := log.New(stderr, "", log.Lshortfile)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	closeTime, err := time.ParseDuration(fmt.Sprintf("%ss", os.Getenv("COFFEE_SHOP_CLOSE_TIME")))
	if err != nil {
		logger.Println("Bad COFFEE_SHOP_CLOSE_TIME")
		exit(1)
		return
	}
	shutdown, err := time.ParseDuration(fmt.Sprintf("%ss", os.Getenv("COFFEE_SHOP_SHUTDOWN")))
	if err != nil {
		logger.Println("Bad COFFEE_SHOP_SHUTDOWN")
		exit(1)
		return
	}
	customers, err := strconv.Atoi(os.Getenv("COFFEE_SHOP_CUSTOMERS"))
	if err != nil {
		logger.Println("Bad COFFEE_SHOP_CUSTOMERS")
		exit(1)
		return
	}
	baristas, err := strconv.Atoi(os.Getenv("COFFEE_SHOP_BARISTAS"))
	if err != nil {
		logger.Println("Bad COFFEE_SHOP_BARISTAS")
		exit(1)
		return
	}
	s := NewStore(logger)
	s.CloseAfter(closeTime)
	s.Customers(RandomGroupOfCustomers(customers))
	s.Baristas(RandomGroupOfBaristas(baristas))
	complete, err := s.Open(ctx)
	if err != nil {
		exit(1)
		return
	}
	select {
	case <-stop:
		cancel()
		select {
		case <-time.After(shutdown):
		case <-complete:
		}
	case <-complete:
	}
}
