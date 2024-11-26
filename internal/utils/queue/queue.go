package queue

import (
	"fmt"
	"sync"
)

type Message interface{}

type worker interface {
	Process(message Message) error
}

// ProcessQueue - очередь для элементов, реализующих интерфейс Message
type ProcessQueue struct {
	data   chan Message
	wg     sync.WaitGroup
	worker worker
}

// NewProcessQueue создает новую очередь с буфером
func NewProcessQueue(bufferSize int, worker worker) *ProcessQueue {
	q := &ProcessQueue{
		data:   make(chan Message, bufferSize),
		worker: worker,
	}
	q.wg.Add(1)

	go q.processQueue()
	return q
}

func (q *ProcessQueue) processQueue() {
	defer q.wg.Done()

	for item := range q.data {
		err := q.worker.Process(item)
		if err != nil {
			fmt.Printf("Ошибка при обработке элемента: %v\n", err)
		}
	}
}

func (q *ProcessQueue) Add(item Message) {
	q.data <- item
}

func (q *ProcessQueue) Close() {
	close(q.data)
	q.wg.Wait()
}
