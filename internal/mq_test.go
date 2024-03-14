package internal

import (
	"fmt"
	"testing"
	"time"
)



func TestNewMessageQueue(t *testing.T) {
	mq := NewMessageQueue()
	if mq == nil {
		t.Error("Expected new MessageQueue, got nil")
	}
}

func TestAddProcess(t *testing.T) {

	mq := NewMessageQueue()

	process := &mockProcess{
		Url: "https://www.youtube.com/watch?v=1234567890",
	}
	go mq.AddProcess(process)
	proc := <-mq.processChannel
	if proc != process {
		t.Error("Expected process in channel, got different process")
	}
}

func TestSubscriber(t *testing.T) {
	mq := NewMessageQueue()
	process := &mockProcess{} // Replace with a real Process instance
	go mq.AddProcess(process)

	go mq.Subscriber()

	time.Sleep(time.Duration(2) * time.Second)
	if process.Progress.Status == StatusCompleted {
		t.Error("Expected process to be inprogress, but got completed")
	}
	time.Sleep(time.Duration(2) * time.Second)
	if process.Progress.Status == StatusCompleted {
		fmt.Printf("Process completed\n")
	} else {
		t.Error("Expected process to be completed, but got inprogress")
	}

}
