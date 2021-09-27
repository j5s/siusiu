package bolt

import "github.com/asdine/storm"

var db *storm.DB

//Init 初始化
func Init() (err error) {
	db, err = storm.Open("biu.db")
	return err
}

//Close 关闭
func Close() {
	db.Close()
}
