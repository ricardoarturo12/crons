package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	gomail "gopkg.in/gomail.v2"
)

var (
	serverHost      string
	serverPort      int
	checkInterval   int
	checkTimeout    int
	printOnlyErrors bool
	errorLogger     *log.Logger
	defaultLogger   *log.Logger
	errorPing       bool
)

func sendEmail() {
	msg := gomail.NewMessage()
	msg.SetHeader("From", os.Getenv("FROM"))
	msg.SetHeader("To", os.Getenv("TO"))
	msg.SetHeader("Subject", os.Getenv("SUBJECT"))
	msg.SetBody("text/html", "<b>Servidor caido</b>")
	// msg.Attach("/home/User/cat.jpg")

	n := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL"), os.Getenv("PASSWORD"))

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}
}

func init() {
	// script arguments
	flag.StringVar(&serverHost, "host", os.Getenv("SERVER"), "Server ip or name to check.")
	sv, _ := strconv.Atoi(os.Getenv("PORT"))
	flag.IntVar(&serverPort, "port", sv, "Server TCP port to check.")
	flag.IntVar(&checkInterval, "interval", 5, "Check interval in seconds.")
	flag.IntVar(&checkTimeout, "timeout", 5, "Connection timeout in seconds.")
	flag.BoolVar(&printOnlyErrors, "only_errors", false, "Print only fails.")
	flag.Parse()

	// loggers
	errorLogger = log.New(os.Stderr, "ERROR: ", log.LstdFlags)
	defaultLogger = log.New(os.Stdout, "", log.LstdFlags)
	if printOnlyErrors {
		defaultLogger.SetOutput(ioutil.Discard)
	}

	// check for argument values
	if serverHost == "" {
		errorLogger.Fatalln("Server host not set.")
	}

	if serverPort == 0 {
		errorLogger.Fatalln("Server port not set.")
	}

}

func main() {
	serverAddress := net.JoinHostPort(serverHost, strconv.Itoa(serverPort))
	timeout := time.Second * time.Duration(checkTimeout)
	defaultLogger.Printf("Starting tcp port check: %s\n", serverAddress)
	for {
		_, tcpErr := net.DialTimeout("tcp", serverAddress, timeout)

		tcpResult := "FAIL"
		if tcpErr == nil {
			tcpResult = "OK"
		}
		if tcpResult == "OK" {
			defaultLogger.Printf("Connection success to \"%s\"\n", serverAddress)
			errorPing = false
		} else {
			if !errorPing {
				errorLogger.Println("servidor caido")
				sendEmail()
			}
			errorPing = true
			errorLogger.Println("servidor sigue caido")
		}
		time.Sleep(time.Second * time.Duration(checkInterval))
	}

}
