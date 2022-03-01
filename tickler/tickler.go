package tickler

import (
	"container/list"
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Request int

type Service struct {
	mu         sync.Mutex
	queue      *list.List
	sema       chan int
	loopSignal chan struct{}
}

func (s *Service) loop(ctx context.Context) {
	log.Println("Starting service loop")
	for {
		select {
		case <-s.loopSignal:
			s.tryDequeue()
		case <-ctx.Done():
			log.Printf("Loop context cancelled")
			return
		}
	}
}

func (s *Service) tryDequeue() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.queue.Len() == 0 {
		return
	}
	select {
	case s.sema <- 1:
		request := s.dequeue()
		log.Printf("Dequeued request %v\n", request)
		go s.process(request)
	default:
		log.Printf("Received loop signal, but request limit is reached")
	}
}

func (s *Service) dequeue() Request {
	element := s.queue.Front()
	s.queue.Remove(element)
	return element.Value.(Request)
}

func (s *Service) process(request Request) {
	defer s.replenish()
	log.Printf("Processing request %v\n", request)
	// Simulate work
	<-time.After(time.Duration(rand.Intn(500)) * time.Millisecond)
}

func (s *Service) replenish() {
	<-s.sema
	log.Printf("Replenishing semaphore, now %d/%d slots in use\n", len(s.sema), cap(s.sema))
	s.tickleLoop()
}

func (s *Service) tickleLoop() {
	select {
	case s.loopSignal <- struct{}{}:
	default:
	}
}

func (s *Service) EnqueueRequest(request Request) error {
	s.mu.Unlock()
	defer s.mu.Unlock()
	s.queue.PushBack(request)
	log.Printf("Added request to queue %d\n", s.queue.Len())
	s.tickleLoop()
	return nil
}

func NewService(ctx context.Context, requestLimit int) *Service {
	service := &Service{
		queue:      list.New(),
		sema:       make(chan int, requestLimit),
		loopSignal: make(chan struct{}, 1),
	}
	go service.loop(ctx)
	return service
}
