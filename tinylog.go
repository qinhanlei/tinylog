package tinylog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const MAX_FILE_LOG_LINE = 100000 // maximum log lines in one file

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

// http://man7.org/linux/man-pages/man5/terminal-colors.d.5.html
const (
	BLACK         = 30
	RED           = 31
	GREEN         = 32
	YELLOW        = 33
	BLUE          = 34
	PURPLE        = 35
	CYAN          = 36
	GRAY          = 37
	COLOR_FORMAT  = "\x1b[%dm"
	COLOR_DEFAULT = "\x1b[m"
)

func fc(c int) string { return fmt.Sprintf(COLOR_FORMAT, c) } // format color
var _lvtag = [...]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
var _lvcolor = [...]string{COLOR_DEFAULT, fc(GREEN), fc(YELLOW), fc(RED), fc(PURPLE)}

var _proc string

var _fidx int // log file index
var _file *os.File
var _flcnt int // log file line count

var _logdir string = "." // the directory of log files, no slash at last
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

func logfname() string {
	return fmt.Sprintf("%s/%s_%s_%02d.log", _logdir, _proc, time.Now().Format("20060102_150405"), _fidx)
}

func logit(lv int, format interface{}, v ...interface{}) {
	if lv < _loglv {
		return
	}
	if _file == nil {
		out := fmt.Sprintf("%s/%s_%s_%02d.log", _logdir, _proc, time.Now().Format("20060102_150405"), _fidx)
		if f, e := os.OpenFile(out, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660); e != nil {
			fmt.Printf("can't open log file:%v error:%v", out, e)
		} else {
			_file = f
			_logger = log.New(_file, "", _logflag)
			_stdlog = log.New(os.Stdout, "", _logflag)
		}
	}
	tag := "[" + _lvtag[lv] + "] " + _proc + " "
	_logger.SetPrefix(tag)
	_stdlog.SetPrefix(_lvcolor[lv] + tag)
	switch format := format.(type) {
	case string:
		_stdlog.Output(3, fmt.Sprintf(format, v...)+_lvcolor[0])
		_logger.Output(3, fmt.Sprintf(format, v...))
	default:
		_stdlog.Output(3, fmt.Sprintf("%v", format)+_lvcolor[0])
		_logger.Output(3, fmt.Sprintf("%v", format))
	}
	_flcnt++
	if _flcnt >= MAX_FILE_LOG_LINE {
		_fidx++
		_flcnt = 0
		Close()
	}
}

func SetDir(logdir string) {
	s, err := os.Stat(logdir)
	if err != nil {
		panic(fmt.Sprintf("get dir:[%v] stat error:%v", logdir, err))
	}
	if !s.IsDir() {
		panic(fmt.Sprintf("get [%v] is not dir!", logdir))
	}
	_logdir = logdir
}
func SetLv(lv int)     { _loglv = lv }
func SetFlag(flag int) { _logflag = flag }
func Close()           { _file.Close(); _file = nil }

func Debug(format interface{}, v ...interface{}) { logit(DEBUG, format, v...) }
func Info(format interface{}, v ...interface{})  { logit(INFO, format, v...) }
func Warn(format interface{}, v ...interface{})  { logit(WARN, format, v...) }
func Error(format interface{}, v ...interface{}) { logit(ERROR, format, v...) }
func Fatal(format interface{}, v ...interface{}) { logit(FATAL, format, v...); os.Exit(-1) }
