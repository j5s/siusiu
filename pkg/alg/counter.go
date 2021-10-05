package alg

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Progress struct {
	total        int64
	finished     int64
	finishedLock *sync.Mutex
	totalLock    *sync.Mutex
}

func NewProgress() *Progress {
	return &Progress{
		finishedLock: new(sync.Mutex),
		totalLock:    new(sync.Mutex),
	}
}

func (p *Progress) AddTotal() {
	p.totalLock.Lock()
	p.total++
	p.totalLock.Unlock()
}

func (p *Progress) AddFinished() {
	p.finishedLock.Lock()
	p.finished++
	p.finishedLock.Unlock()
}

//Show 展示进度
func (p *Progress) Show(ctx context.Context) {
	go func() {
	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case <-time.Tick(time.Second):
				percent := float64(p.finished) / float64(p.total) * 100
				fmt.Printf("\rtotal:%10d finished:%10d percent:%10.2f%%", p.total, p.finished, percent)
			}
		}
	}()
}
