package main

import (
	"log"

	"github.com/arowden/unit_testing/1-interfaces/problems/1/mocks"
	"github.com/arowden/unit_testing/1-interfaces/problems/1/queue"
)

// Given the following consumer. Define the smallest possible interface that could replace the
// sqs.Client and assign it to the queue field of the Consumer.  If done correctly, 'go build'
// should not return an error.
//
// Remove the queue import after assigning the interface.

type Consumer struct {
	queue *queue.Queue
}

func (c *Consumer) Run() {
	for {
		msg, err := c.queue.Peek()
		if err != nil {
			continue
		}

		err = process(msg)
		if err != nil {
			continue
		}

		err = c.queue.Delete(msg)
		if err != nil {
			log.Println(err)
		}
	}
}

func process(msg string) error {
	return nil
}

func main() {
	mockClient := &mocks.Queue{}
	_ = Consumer{queue: mockClient}
}
