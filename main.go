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

// nolint
var (
	stop             = make(chan os.Signal, 1)
	exit             = os.Exit
	stderr io.Writer = os.Stderr
)

func main() {
	logger := log.New(stderr, "", log.Lshortfile)
	ctx, cancel := context.WithCancel(context.Background())
	errorCnt := 0

	defer cancel()

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	closeTime, err := time.ParseDuration(fmt.Sprintf("%ss", os.Getenv("COFFEE_SHOP_CLOSE_TIME")))
	if err != nil {
		logger.Println("please specify number, in seconds, for envvar COFFEE_SHOP_CLOSE_TIME")
		errorCnt++
	}

	shutdown, err := time.ParseDuration(fmt.Sprintf("%ss", os.Getenv("COFFEE_SHOP_SHUTDOWN")))
	if err != nil {
		logger.Println("please specify number, in seconds, for envvar COFFEE_SHOP_SHUTDOWN")
		errorCnt++
	}

	customers, err := strconv.Atoi(os.Getenv("COFFEE_SHOP_CUSTOMERS"))
	if err != nil {
		logger.Println("please specify number of customers for envvar COFFEE_SHOP_CUSTOMERS")
		errorCnt++
	}

	baristas, err := strconv.Atoi(os.Getenv("COFFEE_SHOP_BARISTAS"))
	if err != nil {
		logger.Println("please specify number of baristas for envvar COFFEE_SHOP_BARISTAS")
	}

	if errorCnt != 0 {
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

	time.Sleep(1 * time.Second)
	select {
	case <-stop:
		logger.Println("I received a signal to close the store")
		cancel()
		select {
		case <-time.After(shutdown):
			logger.Println("Shutdown time reached, closing anyway")
			<-complete
		case <-complete:
			logger.Println("Store closed")
		}
	case <-complete:
		logger.Println("Store closed")
	}
	exit(0)
}
