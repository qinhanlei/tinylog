Tiny simple naive [![Build Status](https://travis-ci.org/qinhanlei/tinylog.svg?branch=master)](https://travis-ci.org/qinhanlei/tinylog)

## API
```golang
func Init(logdir string) {} // *.log files DIR path required
func SetLv(lv int)       {} // lv in (DEBUG = iota, INFO, WARN, ERROR, FATAL)
func SetFlag(flag int)   {} // see standard pkg log.LstdFlags
func Close()             {}
// utils
func Debug(format interface{}, v ...interface{}) {}
func Info(format interface{}, v ...interface{})  {}
func Warn(format interface{}, v ...interface{})  {}
func Error(format interface{}, v ...interface{}) {}
func Fatal(format interface{}, v ...interface{}) {}
```

## Usage
`go get github.com/qinhanlei/tinylog`
```golang
package main

import log "github.com/qinhanlei/tinylog"

func init() { log.Init(".") }

func main() {
	log.Debug("Hello Debug")
	log.Info("Hello Info")
	log.Warn("Hello Warn")
	log.Error("Hello Error")
	log.Fatal("Hello Fatal")
}
```
<div align="left"><img width="500" height="90" src="https://raw.githubusercontent.com/qinhanlei/tinylog/master/test.png"/></div>
