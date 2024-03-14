package internal

import (
	"github.com/hash167/yt-dl-rpc/server/config"
)

type MessageQueue struct {
	processChannel chan ProcessInterface
}

// Create a new message queue, default size is equal to the number of logical CPUs
func NewMessageQueue() *MessageQueue {
	return &MessageQueue{
		processChannel: make(chan ProcessInterface, config.Instance().QueueSize),
	}
}

// Add a new process to the channel
func (mq *MessageQueue) AddProcess(process ProcessInterface) {
	mq.processChannel <- process
}

// Subscribe to the channel to get new processes
func (mq MessageQueue) Subscriber() {
	for process := range mq.processChannel {
		go process.Start()
	}
}
