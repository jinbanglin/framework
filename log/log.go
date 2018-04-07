package log

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

var gLogger *Logger

func init() {
	chaos()
	if out == "file" {
		go poller()
	}
}

func chaos() {
	loadConfig()
	y, m, d := time.Now().Date()
	if gLogger == nil {
		gLogger = &Logger{
			look:        uint32(coreDead),
			fileName:    fileName,
			fileBufSize: bufSize,
			path:        filepath.Join(filePath, fileName),
			timestamp:   y*10000 + int(m)*100 + d*1,
			fileMaxSize: maxSize,
			bucket:      make(chan *bytes.Buffer, bucketLen),
			closeSignal: make(chan string),
			lock:        &sync.RWMutex{},
		}
	}
}

func (l *Logger) loadCurLogFile() error {
	l.link = filepath.Join(l.path, fileName+".log")
	actFileName, ok := isLinkFile(l.link)
	if !ok {
		return errors.New("is not link file")
	}
	l.fileName = filepath.Join(l.path, actFileName)
	f, err := openFile(l.fileName)
	if err != nil {
		return err
	}
	info, err := os.Stat(l.fileName)
	if err != nil {
		return err
	}
	sp := strings.Split(actFileName, ".")
	t, err := time.Parse("2006-01-02", sp[1])
	if err != nil {
		fmt.Errorf("loadCurrentLogFile |err=%v", err)
		return err
	}
	y, m, d := t.Date()
	l.timestamp = y*10000 + int(m)*100 + d*1
	l.file = f
	l.fileActualSize = int(info.Size())
	l.fileWriter = bufio.NewWriterSize(f, l.fileBufSize)
	return nil
}

func (l *Logger) createFile() (err error) {
	if !pathIsExist(l.path) {
		if err = os.MkdirAll(l.path, os.ModePerm); err != nil {
			return
		}
	}
	now := time.Now()
	y, m, d := now.Date()
	l.timestamp = y*10000 + int(m)*100 + d*1
	l.fileName = filepath.Join(
		l.path,
		filepath.Base(os.Args[0])+"."+now.Format("2006-01-02.15.04.05.000")+".log")
	f, err := openFile(l.fileName)
	if err != nil {
		return err
	}
	l.file = f
	l.fileActualSize = 0
	l.fileWriter = bufio.NewWriterSize(f, l.fileBufSize)
	l.link = filepath.Join(l.path, fileName+".log")
	return createLinkFile(l.fileName, l.link)
}

func (l *Logger) sync() {
	if l.lookRunning() {
		l.fileWriter.Flush()
	}
}

const fileMaxDelta = 100

func (l *Logger) rotate() bool {
	if !l.lookRunning() {
		return false
	}
	y, m, d := time.Now().Date()
	timestamp := y*10000 + int(m)*100 + d*1
	if l.fileActualSize <= l.fileMaxSize-fileMaxDelta && timestamp <= l.timestamp {
		return false
	}
	l.fileWriter.Flush()
	closeFile(l.file)
	if err := l.createFile(); err != nil {
		return false
	}
	return true
}

func (l *Logger) lookRunning() bool {
	if atomic.LoadUint32(&l.look) == uint32(coreRunning) {
		return true
	}
	return false
}

func (l *Logger) lookDead() bool {
	if atomic.LoadUint32(&l.look) == uint32(coreDead) {
		return true
	}
	return false
}

func (l *Logger) lookBlock() bool {
	if atomic.LoadUint32(&l.look) == uint32(coreBlock) {
		return true
	}
	return false
}

func (l *Logger) signalHandler() {
	var sigChan = make(chan os.Signal)
	//syscall.SIGHUP
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT)
	for {
		select {
		case sig := <-sigChan:
			l.closeSignal <- "close"
			fmt.Println("LOGX receive os signal is ", sig)
			l.fileWriter.Flush()
			closeFile(l.file)
			atomic.SwapUint32(&l.look, uint32(coreDead))
			close(l.bucket)
			os.Exit(1)
		}
	}
}

func (l *Logger) release(buf *bytes.Buffer) { bufferPoolFree(buf) }

func caller() string {
	if pc, f, l, ok := runtime.Caller(2); ok {
		funcName := runtime.FuncForPC(pc).Name()
		return path.Base(f) + "|" + path.Base(funcName) + "|" + strconv.Itoa(l)
	}
	//pc := make([]uintptr, 3, 3)
	//cnt := runtime.Callers(6, pc)
	//
	//for i := 0; i < cnt; i++ {
	//	fu := runtime.FuncForPC(pc[i] - 1)
	//	name := fu.Name()
	//
	//	if !strings.Contains(name, "github.com/kafrax/logx") {
	//		f, l := fu.FileLine(pc[i] - 1)
	//		return path.Base(f) + "|" + path.Base(name) + "|" + strconv.Itoa(l)
	//	}
	//
	//	if pc, f, l, ok := runtime.Caller(8); ok {
	//		funcName := runtime.FuncForPC(pc).Name()
	//		return path.Base(f) + "|" + path.Base(funcName) + "|" + strconv.Itoa(l)
	//	}
	//}
	return ""
}

func print(buf *bytes.Buffer) {
	switch out {
	case "file":
		gLogger.bucket <- buf
	case "stdout":
		fmt.Print(buf.String())
	case "kafka": //todo send to kafka etc.
	case "nsq": //todo send to kafka nsq etc.
	default:
		fmt.Println(buf.String())
	}
}

func Debugf(format string, msg ... interface{}) {
	if levelFlag > _DEBUG {
		return
	}
	buf := bufferPoolGet()
	buf.Write(string2bytes("[DEBU][" + time.Now().Format("01-02.15.04.05.000") + "]" + "[" + caller() + "] "))
	buf.Write(string2bytes(fmt.Sprintf(format, msg...) + "\n"))
	print(buf)
}

func Infof(format string, msg ... interface{}) {
	if levelFlag > _INFO {
		return
	}
	buf := bufferPoolGet()
	buf.Write(string2bytes("[INFO][" + time.Now().Format("01-02.15.04.05.000") + "]" + "[" + caller() + "] "))
	buf.Write(string2bytes(fmt.Sprintf(format, msg...) + "\n"))
	print(buf)
}

func Warnf(format string, msg ... interface{}) {
	if levelFlag > _WARN {
		return
	}
	buf := bufferPoolGet()
	buf.Write(string2bytes("[WARN][" + time.Now().Format("01-02.15.04.05.000") + "]" + "[" + caller() + "] "))
	buf.Write(string2bytes(fmt.Sprintf(format, msg...) + "\n"))
	print(buf)
}

func Errorf(format string, msg ... interface{}) {
	if levelFlag > _ERR {
		return
	}
	buf := bufferPoolGet()
	buf.Write(string2bytes("[ERRO][" + time.Now().Format("01-02.15.04.05.000") + "]" + "[" + caller() + "] "))
	buf.Write(string2bytes(fmt.Sprintf(format, msg...) + "\n"))
	print(buf)
}

func Fatalf(format string, msg ... interface{}) {
	if levelFlag > _DISASTER {
		return
	}
	buf := bufferPoolGet()
	buf.Write(string2bytes("[FTAL][" + time.Now().Format("01-02.15.04.05.000") + "]" + "[" + caller() + "] "))
	buf.Write(string2bytes(fmt.Sprintf(format, msg...) + "\n"))
	print(buf)
}

func Stackf(format string, msg ... interface{}) {
	s := fmt.Sprintf(format, msg...)
	s += "\n"
	buf := make([]byte, 1<<20)
	n := runtime.Stack(buf, true)
	s += string(buf[:n])
	s += "\n"
	fmt.Println("[STAC][" + time.Now().Format("01-02.15.04.05.000") + "]" + "[" + caller() + "] " + s)
}

func Debug(msg ... interface{}) {
	if levelFlag > _DEBUG {
		return
	}
	buf := bufferPoolGet()
	buf.Write(string2bytes("[DEBU][" + time.Now().Format("01-02.15.04.05.000") + "]" + "[" + caller() + "] "))
	buf.Write(string2bytes(fmt.Sprintln(msg...)))
	print(buf)
}

func Info(msg ... interface{}) {
	if levelFlag > _INFO {
		return
	}
	buf := bufferPoolGet()
	buf.Write(string2bytes("[INFO][" + time.Now().Format("01-02.15.04.05.000") + "]" + "[" + caller() + "] "))
	buf.Write(string2bytes(fmt.Sprintln(msg...)))
	print(buf)
}

func Warn(msg ... interface{}) {
	if levelFlag > _WARN {
		return
	}
	buf := bufferPoolGet()
	buf.Write(string2bytes("[WARN][" + time.Now().Format("01-02.15.04.05.000") + "]" + "[" + caller() + "] "))
	buf.Write(string2bytes(fmt.Sprintln(msg...)))
	print(buf)
}

func Error(msg ... interface{}) {
	if levelFlag > _ERR {
		return
	}
	buf := bufferPoolGet()
	buf.Write(string2bytes("[ERRO][" + time.Now().Format("01-02.15.04.05.000") + "]" + "[" + caller() + "] "))
	buf.Write(string2bytes(fmt.Sprintln(msg...)))
	print(buf)
}

func Fatal(msg ... interface{}) {
	if levelFlag > _DISASTER {
		return
	}
	buf := bufferPoolGet()
	buf.Write(string2bytes("[FTAL][" + time.Now().Format("01-02.15.04.05.000") + "]" + "[" + caller() + "] "))
	buf.Write(string2bytes(fmt.Sprintln(msg...)))
	print(buf)
}

func Stack(msg ... interface{}) {
	s := fmt.Sprintln(msg...)
	s += "\n"
	buf := make([]byte, 1<<20)
	n := runtime.Stack(buf, true)
	s += string(buf[:n])
	s += "\n"
	fmt.Println("[STAC][" + time.Now().Format("01-02.15.04.05.000") + "]" + "[" + caller() + "] " + s)
}
