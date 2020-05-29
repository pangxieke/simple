package log

import (
	"io"
	"log"
	"os"
	"path"
	"runtime"
)

var (
	logInfo *log.Logger
)

func init() {
	errFile, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}

	logInfo = log.New(io.MultiWriter(os.Stderr, errFile), "Info:", log.Ldate)

}

func Info(args ...interface{}) {
	filename, line := callFile()

	logInfo.Println(filename, line, args)
}

func Fatal(args ...interface{}) {
	filename, line := callFile()
	log.Fatal(filename, line, args)
}

// 错误位置
func callFile() (string, int) {
	callDepth := 1
	_, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		file = "???"
		line = 0
	}
	_, filename := path.Split(file)
	return filename, line
}
