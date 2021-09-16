package scanner

import (
	"biu/logger"
	"biu/models"
	"fmt"
	"runtime"
	"sync"

	"gopkg.in/cheggaaa/pb.v2"
)

//PasswdGuesser 猜解器
type PasswdGuesser struct {
	wg           *sync.WaitGroup
	routineCount int
	taskChans    []chan models.Task
	resultCh     chan models.Task
	foundMap     map[string]string
	progressBar  *pb.ProgressBar
	passwdList   []string
	userList     []string
	targetList   []models.TargetWithHandler
}

//NewPasswdGuesser 猜解器的构造函数
func NewPasswdGuesser(passwdList, userList []string, targetList []models.TargetWithHandler, routineCount int) *PasswdGuesser {
	var wg sync.WaitGroup
	//已经完成的扫描任务
	foundmap := make(map[string]string)
	//任务管道
	taskChans := make([]chan models.Task, 0, routineCount)
	//结果管道
	resultCh := make(chan models.Task, routineCount*10)
	//计算总任务数初始化进度条
	totalTaskCount := len(passwdList) * len(userList) * len(targetList)
	progressBar := pb.StartNew(totalTaskCount)
	progressBar.SetTemplate(`{{ rndcolor "Scanning progress: " }} {{  percent . "[%.02f%%]" "[?]"| rndcolor}} {{ counters . "[%s/%s]" "[%s/?]" | rndcolor}} {{ bar . "「" "-" (rnd "ᗧ" "◔" "◕" "◷" ) "•" "」" | rndcolor }} {{rtime . | rndcolor}} `)
	return &PasswdGuesser{
		taskChans:    taskChans,
		resultCh:     resultCh,
		routineCount: routineCount,
		wg:           &wg,
		foundMap:     foundmap,
		progressBar:  progressBar,
		passwdList:   passwdList,
		userList:     userList,
		targetList:   targetList,
	}
}

//GenTask 生成任务
func (g *PasswdGuesser) GenTask(passwdList, userList []string, targetList []models.TargetWithHandler) {
	var i int = 0
	for passwdIndex := range passwdList {
		for userIndex := range userList {
			for targetIndex := range targetList {
				//如果任务已完成，停止生成同类任务
				_, ok := g.foundMap[targetList[targetIndex].Serialize()]
				if ok {
					logger.Log.Debugf("target:%s skipped", targetList[targetIndex].Serialize())
					continue
				}
				task := models.Task{
					TargetWithHandler: targetList[targetIndex],
					Credential: models.Credential{
						Username: userList[userIndex],
						Password: passwdList[passwdIndex],
					},
				}
				g.taskChans[i] <- task
				i = (i + 1) % g.routineCount
			}
		}
	}
	//生产完毕，关闭所有管道的写通道
	for _, taskCh := range g.taskChans {
		close(taskCh)
	}
	//阻塞等待消费者协程结束工作
	g.wg.Wait()
	// logger.Log.Debugf("所有消费者协程工作完成")
	//关闭消费协程的输出管道
	g.wg.Add(1)
	close(g.resultCh)
	g.wg.Wait()
}

//DealTask 消费任务
func (g *PasswdGuesser) DealTask() {
	g.wg.Add(g.routineCount)
	for i := 0; i < g.routineCount; i++ {
		c := make(chan models.Task, g.routineCount*2)
		go func() {
			for task := range c {
				//1.如果任务已完成,停止执行同类任务
				_, ok := g.foundMap[task.Target.Serialize()]
				if ok {
					logger.Log.Debugf("[skip]target:%s skipped,routine:%d", task.Target.Serialize(), runtime.NumGoroutine())
					continue
				}
				logger.Log.Infof("[*]checking: %v:%d|%s, UserName: %v, Password: %v, goroutineNum: %v", task.IP, task.Port,
					task.Protocal, task.Username, task.Password, runtime.NumGoroutine())
				success, err := task.Check()
				g.progressBar.Increment()
				if err != nil {
					continue
				}
				if success {
					g.resultCh <- task
				}
			}
			g.wg.Done()
			// logger.Log.Debugf("routine:%d done", runtime.NumGoroutine())
		}()
		g.taskChans = append(g.taskChans, c)
	}
}

//PrintResult 输出结果
func (g *PasswdGuesser) PrintResult() {
	go func() {
		for task := range g.resultCh {
			g.foundMap[task.Target.Serialize()] = task.Credential.Serialize()
			logger.Log.Printf("[Found]:%s", task.Serialize())
		}
		g.wg.Done()
	}()

}

//Summary 总结扫描结果
func (g *PasswdGuesser) Summary() {
	g.progressBar.Finish()
	fmt.Printf("[*]扫描完成,成功数:%d\n", len(g.foundMap))
	for key, value := range g.foundMap {
		fmt.Println(key, value)
	}
}
