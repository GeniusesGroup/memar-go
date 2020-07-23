/* For license and copyright information please see LEGAL file in repository */

package log

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

// TODO::: fix problem with multi CPU core parallelism

// hold logs until app running.
var buffer = make([]byte, 0, 8192)

// Debug show log in standard console & append log to buffer to save them later.
func Debug(a ...interface{}) {
	var log = fmt.Sprintln("[Debug]", time.Now().Format(timeFormat), a)
	// write log to stderr
	os.Stderr.WriteString(log)
	// Append log to buffer for saving by SaveToStorage() later.
	buffer = append(buffer, log...)
}

// Info show log in standard console & append log to buffer to save them later.
func Info(a ...interface{}) {
	var log = fmt.Sprintln("[Info]", time.Now().Format(timeFormat), a)
	// write log to stderr
	os.Stderr.WriteString(log)
	// Append log to buffer for saving by SaveToStorage() later.
	buffer = append(buffer, log...)
}

// Warn show log in standard console & append log to buffer to save them later.
func Warn(a ...interface{}) {
	var log = fmt.Sprintln("[Warn]", time.Now().Format(timeFormat), a)
	// write log to stderr
	os.Stderr.WriteString(log)
	// Append log to buffer for saving by SaveToStorage() later.
	buffer = append(buffer, log...)
}

// Fatal show log in standard console & append log to buffer to save them later and exit app.
func Fatal(a ...interface{}) {
	var log = fmt.Sprintln("[Fatal]", time.Now().Format(timeFormat), a)
	// write log to stderr
	os.Stderr.WriteString(log)
	// Append log to buffer for saving by SaveToStorage() later.
	buffer = append(buffer, log...)
	panic("Due to important log, panic situation occur")
}

// SaveToStorage use to make||flush to {location}/{name}.log if needed!
func SaveToStorage(name, location string) {
	// TODO::: Check if log file not exist create it || Flush old logs!
	err := ioutil.WriteFile(path.Join(location, name+".log"), buffer, 0755)
	if err != nil {
		panic(fmt.Sprintf("Unable to write log file: %v", err))
	}
}
