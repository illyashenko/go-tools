package concurrent

type Semaphore struct {
	ch chan struct{}
}

func InitSemaphore(max int) *Semaphore {
	return &Semaphore{
		ch: make(chan struct{}, max),
	}
}

func (s *Semaphore) Acquire() {
	s.ch <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.ch
}
