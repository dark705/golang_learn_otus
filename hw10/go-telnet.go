package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	if len(os.Args) < 3 {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", "Please set host and port: example go-telnet host port [--timeout=10s] \n --timeout=10s - optional, default = 10s")
		os.Exit(2)
	}
	host := os.Args[1]
	port := os.Args[2]

	var timeout int
	flag.IntVar(&timeout, "timeout", 10, "time out in seconds")
	flag.Parse()
	connect, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), time.Second*time.Duration(timeout))
	failOnError("Can't connect to remote host", err)
	ctx, chancel := context.WithCancel(context.Background())

	chancelCh := make(chan struct{}, 1)
	go readerFromStdIn(ctx, connect, chancelCh)
	go writerToStdIn(ctx, connect, chancelCh)
	waitForNeedShutdown(connect, chancel, chancelCh)
}

func waitForNeedShutdown(conn net.Conn, chancel context.CancelFunc, ch chan struct{}) {
	defer func() {
		fmt.Println("Close connection...")
		chancel()
		err := conn.Close()
		failOnError("Fail on close connection", err)
		close(ch)
		fmt.Println("Exit.")
	}()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	for {
		select {
		case <-ch:
			return
		case v := <-osSignals:
			fmt.Printf("Get signal from os: %v\n", v)
			return
		}
	}
}

func readerFromStdIn(ctx context.Context, connect net.Conn, ch chan struct{}) {
	defer func() {
		ch <- struct{}{}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if !scanner.Scan() {
				return
			}
			str := scanner.Text()
			_, err := connect.Write([]byte(fmt.Sprintf("%s\n", str)))
			if err != nil {
				return
			}
		}
	}
}

func writerToStdIn(ctx context.Context, connect net.Conn, ch chan struct{}) {
	defer func() {
		ch <- struct{}{}
	}()
	scanner := bufio.NewScanner(connect)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if !scanner.Scan() {
				return
			}
			text := scanner.Text()
			fmt.Println(text)
		}
	}
}

func failOnError(message string, err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s: %v\n", message, err)
		os.Exit(2)
	}
}
