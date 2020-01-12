package main

import (
	"file/reader/txt"
	"fmt"
	"os"
	"time"
)

func main() {

	file := txt.New("file_to_read.txt")

	ok, err := file.Read()

	if !ok {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	contentToSend := file.AsStringSlice()

	t := time.NewTimer(5 * time.Second)
	messagesCh := make(chan string)
	printCh := make(chan string)

	sender(contentToSend, messagesCh)

	for i := 0; i < 5; i++ {
		go receiver(i, messagesCh, printCh)
	}

	for {

		select {
		case print := <-printCh:
			fmt.Printf("%s", print)
		case <-t.C:
			fmt.Println("Mensagend processadas!!")
			close(printCh)
			os.Exit(1)
		}
	}
}

func sender(messagesToSend []string, messagesCh chan string) {
	for _, messageToSend := range messagesToSend {
		go addMessageToChanel(messageToSend, messagesCh)
	}
}

func addMessageToChanel(message string, chanel chan string) {
	chanel <- message
}

func receiver(worker int, messagesCh chan string, printCh chan string) {
	for {
		select {
		case message, ok := <-messagesCh:
			if ok {
				printCh <- fmt.Sprintf("%s worker:%d\n", message, worker)
			} else {
				printCh = nil
			}
		}
	}
}
