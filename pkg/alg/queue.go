package alg

import (
	"fmt"
	"os"
	"sync"
)

//Queue 队列
type Queue struct {
	ch        chan string
	size      int64
	total     int64
	sizeLock  *sync.Mutex
	totalLock *sync.Mutex
}

//NewQueue 构造函数
func NewQueue(bufSize int) (q *Queue) {
	ch := make(chan string, bufSize)
	return &Queue{
		ch:        ch,
		sizeLock:  new(sync.Mutex),
		totalLock: new(sync.Mutex),
	}
}
func (q *Queue) addSize() {
	q.sizeLock.Lock()
	q.size++
	q.sizeLock.Unlock()
}
func (q *Queue) subSize() {
	q.sizeLock.Lock()
	q.size--
	q.sizeLock.Unlock()
}

func (q *Queue) addTotal() {
	q.totalLock.Lock()
	q.total++
	q.totalLock.Unlock()
}

//EnQueue 入队
func (q *Queue) EnQueue(item string) {
	q.ch <- item
	q.addSize()
	q.addTotal()
	q.ShowProgress("入队")
}

//DeQueue 出队
func (q *Queue) DeQueue() (string, bool) {
	item, ok := <-q.ch
	if !ok {
		return item, ok
	}
	q.subSize()
	return item, ok
}

//Close 关闭队列
func (q *Queue) Close() {
	close(q.ch)
}

//GetSize 获取当前队列的元素个数
func (q *Queue) GetSize() (size int64) {
	return q.size
}

//ShowProgress 显示进度
func (q *Queue) ShowProgress(comment string) {
	fmt.Fprintf(os.Stdout, "\r%s total:%d current:%d ", comment, q.total, q.size)
}
