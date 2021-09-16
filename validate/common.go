package validate

import (
	"fmt"
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/sirupsen/logrus"
)

//ReadGoRoutineCount 读取协程数
func ReadGoRoutineCount(c *ishell.Context) (int, error) {
	c.Print(fmt.Sprintf("请输入协程数(默认为%s):", defaultRoutineCount))
	routineCount := c.ReadLine()
	if len(routineCount) == 0 {
		routineCount = defaultRoutineCount
	}
	r, err := strconv.Atoi(routineCount)
	if err != nil {
		logrus.Error("strconv.Atoi(routineCount) failed,err:", err)
		return r, err
	}
	return r, nil
}
