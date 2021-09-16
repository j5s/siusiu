package scanner

import (
	"biu/logger"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
)

//PortScanner 端口扫描器
type PortScanner struct {
	routineCount int
	connectFunc  ConnectFunc
	wg           *sync.WaitGroup
	taskChans    []chan string
}

//NewPortScanner 猜解器的构造函数
func NewPortScanner(routineCount int, mode string) *PortScanner {
	var wg sync.WaitGroup
	taskChans := make([]chan string, 0, routineCount)
	connectFunc, ok := ModeScanFuncMap[mode]
	if !ok {
		logger.Log.Fatalf("mode:%s 不存在", mode)
	}
	return &PortScanner{
		taskChans:    taskChans,
		routineCount: routineCount,
		wg:           &wg,
		connectFunc:  connectFunc,
	}
}

//GenTask 生产扫描任务
func (p *PortScanner) GenTask(hosts []net.IP, ports []int) {
	var i int = 0
	for _, host := range hosts {
		for _, port := range ports {
			p.taskChans[i] <- fmt.Sprintf("%s:%d", host.String(), port)
			i = (i + 1) % p.routineCount
		}
	}
	for _, channel := range p.taskChans {
		close(channel)
	}
	p.wg.Wait()
}

//DealTask 消费任务
func (p *PortScanner) DealTask(count int) {
	for i := 0; i < count; i++ {
		c := make(chan string, p.routineCount*2)
		go func() {
			for target := range c {
				parts := strings.Split(target, ":")
				port, err := strconv.Atoi(parts[1])
				if err != nil {
					log.Println("strconv.Atoi(parts[1]) failed,err:", err)
					continue
				}
				if p.connectFunc(parts[0], port) {
					fmt.Println(target)
				}
			}
			p.wg.Done()
		}()
		p.taskChans = append(p.taskChans, c)
		p.wg.Add(1)
	}
}
