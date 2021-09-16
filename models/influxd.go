package models

import (
	"fmt"
	"time"
)

//LeakInfo 泄漏的敏感信息
// "system": {
// 	"currentTime": "2021-09-15T08:16:28.264082177Z",
// 	"started": "2021-09-15T04:02:02.881384845Z",
// 	"uptime": 15265
// },
// "cmdline": [
// 	"/usr/bin/influxd",
// 	"-config",
// 	"/etc/influxdb/influxdb.conf"
// ],
type LeakInfo struct {
	System    `json:"system"`
	CmdLine   []string `json:"cmdline"`
	Databases []DataBaseInfo
}

//System 系统信息
type System struct {
	CurrentTime time.Time `json:"currentTime"`
	StartedTime time.Time `json:"started"`
	UpTime      int       `json:"uptime"`
}

//DataBaseInfo 数据库信息
type DataBaseInfo struct {
	Name   string `json:"name" mapstructure:"name"`
	Tags   Tags   `json:"tags" mapstructure:"tags"`
	Values Values `json:"value" mapstructure:"values"`
}

//Tags 数据库的标签
type Tags struct {
	Name string `mapstructure:"database"`
}

//Values 数据信息的值
type Values struct {
	NumMeasurments int `json:"numMeasurements" mapstructure:"numMeasurements"`
	NumSeries      int `json:"numSeries" mapstructure:"numSeries"`
}

//Version 软件版本号
// A 表示大版本号，一般当软件整体重写，或出现不向后兼容的改变时，增加A，A为零时表示软件还在开发阶段。
// B 表示功能更新，出现新功能时增加B
// C 表示小修改，如修复bug，只要有修改就增加C
type Version struct {
	A int
	B int
	C int
}

//NewVersion 构造函数
func NewVersion(str string) (*Version, error) {
	v := new(Version)
	_, err := fmt.Sscanf(str, "%d.%d.%d", &v.A, &v.B, &v.C)
	return v, err
}

func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.A, v.B, v.C)
}
