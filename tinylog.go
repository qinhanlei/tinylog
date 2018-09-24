package tinylog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

// http://man7.org/linux/man-pages/man5/terminal-colors.d.5.html
func fc(c int) string { return fmt.Sprintf("\x1b[%dm", c) }

var _lvtag = [...]string{"DEBUG", "INFO ", "WARN ", "ERROR", "FATAL"}
var _lvcolor = [...]string{"\x1b[m", fc(32), fc(33), fc(31), fc(35)}

var _proc string
var _file *os.File
var _loglv = DEBUG
var _logger *log.Logger
var _logflag = log.Ldate | log.Ltime | log.Lshortfile
var _stdlog *log.Logger

func init() {
	proc, ext := filepath.Base(os.Args[0]), filepath.Ext(os.Args[0])
	_proc = string([]byte(proc)[:len(proc)-len(ext)])
	if runtime.GOOS != "linux" && runtime.GOOS != "darwin" {
		_lvcolor = [...]string{"", "", "", "", ""} //NOTE: color not supported yet
	}
}

func logit(lv int, format interface{}, v ...interface{}) {
	if _file != nil && lv >= _loglv {
		tag := "[" + _lvtag[lv] + "] " + _proc + " "
		_logger.SetPrefix(tag)
		_stdlog.SetPrefix(_lvcolor[lv] + tag + _lvcolor[0])
		switch format := format.(type) {
		case string:
			_stdlog.Output(3, fmt.Sprintf(format, v...))
			_logger.Output(3, fmt.Sprintf(format, v...))
		default:
			_stdlog.Output(3, fmt.Sprintf("%v", format))
			_logger.Output(3, fmt.Sprintf("%v", format))
		}
	}
}

func Init(logdir string) {
	if _file == nil {
		var e error
		out := fmt.Sprintf("%s/%s_%s.log", logdir, _proc, time.Now().Format("20060102_150405"))
		_file, e = os.OpenFile(out, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
		if e != nil {
			panic(fmt.Sprintf("can't open log file:%v error:%v", out, e))
		}
		_logger = log.New(_file, "", _logflag)
		_stdlog = log.New(os.Stdout, "", _logflag)
	}
}
func SetLv(lv int)     { _loglv = lv }
func SetFlag(flag int) { _logflag = flag }
func Close()           { _file.Close(); _file = nil }

func Debug(format interface{}, v ...interface{}) { logit(DEBUG, format, v...) }
func Info(format interface{}, v ...interface{})  { logit(INFO, format, v...) }
func Warn(format interface{}, v ...interface{})  { logit(WARN, format, v...) }
func Error(format interface{}, v ...interface{}) { logit(ERROR, format, v...) }
func Fatal(format interface{}, v ...interface{}) { logit(FATAL, format, v...); os.Exit(-1) }
