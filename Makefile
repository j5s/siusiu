.PHONY: all build_linux build_mac build_win run gotool clean help

BINARY="biu"

all: gotool build_mac

build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}-linux
build_win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BINARY}.exe	
build_mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${BINARY}


run:
	@go run ./

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build_linux - 生成linux下的可执行文件"
	@echo "make build_win - 生成windows下的可执行文件"
	@echo "make build_mac - 生成macOS下的可执行文件"
	@echo "make run - 运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"
