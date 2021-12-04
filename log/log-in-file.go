/* For license and copyright information please see LEGAL file in repository */

package log

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	etime "../earth-time"
	"../protocol"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

// hold logs until app running.
// TODO::: check for any problem with multi CPU core parallelism (data race condition)
var logFile *os.File

// Init will initialize log to do some interval saving
func Init(name, repoLocation string, interval etime.Duration) (err error) {
	var logFolder = path.Join(repoLocation, "log")
	os.Mkdir(logFolder, 0700)
	logFile, err = os.OpenFile(path.Join(logFolder, name+strconv.FormatInt(etime.Now().RoundTo(interval), 10)+".log"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0700)
	if err != nil {
		return
	}
	go intervalSaving(name, logFolder, interval)
	return
}

// Debug show log in standard console if requested by protocol.AppScreenMode const & append log to buffer to save them later.
func Debug(a ...interface{}) {
	var log = fmt.Sprintln("[Debug]", time.Now().Format(timeFormat), a)
	if protocol.AppScreenMode {
		os.Stderr.WriteString(log)
	}
	logFile.WriteString(log)
}

// DeepDebug show deep debug log in standard console if requested by protocol.AppScreenMode const & append log to buffer to save them later.
func DeepDebug(a ...interface{}) {
	var log = fmt.Sprintln("[Deep]", time.Now().Format(timeFormat), a)
	if protocol.AppScreenMode {
		os.Stderr.WriteString(log)
	}
	logFile.WriteString(log)
}

// Info show log in standard console if requested by protocol.AppScreenMode const & append log to buffer to save them later.
func Info(a ...interface{}) {
	var log = fmt.Sprintln("[Info]", time.Now().Format(timeFormat), a)
	if protocol.AppScreenMode {
		os.Stderr.WriteString(log)
	}
	logFile.WriteString(log)
}

// Warn show log in standard console if requested by protocol.AppScreenMode const & append log to buffer to save them later.
func Warn(a ...interface{}) {
	var log = fmt.Sprintln("[Warn]", time.Now().Format(timeFormat), a)
	if protocol.AppScreenMode {
		os.Stderr.WriteString(log)
	}
	logFile.WriteString(log)
}

// Panic show log in standard console if requested by protocol.AppScreenMode const & append log to buffer to save them later and throw panic message.
func Panic(a ...interface{}) {
	var log = fmt.Sprintln("[Panic]", time.Now().Format(timeFormat), a)
	if protocol.AppScreenMode {
		os.Stderr.WriteString(log)
	}
	logFile.WriteString(log)
	panic("Due to important log, panic situation occur")
}

// Fatal show log in standard console if requested by protocol.AppScreenMode const & append log to buffer to save them later and exit app.
// Simple use as protocol.App.LogFatal(function()) and not check return error from function
func Fatal(a ...interface{}) {
	if a == nil {
		return
	}
	var log = fmt.Sprintln("[Fatal]", time.Now().Format(timeFormat), a)
	if protocol.AppScreenMode {
		os.Stderr.WriteString(log)
	}
	logFile.WriteString(log)
	os.Exit(125)
}

// SaveToStorage use to make||flush file!
func SaveToStorage() {
	logFile.Close()
}

func intervalSaving(name, location string, interval etime.Duration) {
	var timer = time.NewTimer(time.Duration(etime.Now().UntilRoundTo(interval)) * time.Second)
	for {
		select {
		// case shutdownFeedback := <-pcs.shutdownSignal:
		// 	timer.Stop()
		// 	shutdownFeedback <- struct{}{}
		// 	return
		case <-timer.C:
			logFile.Close()

			logFile, _ = os.OpenFile(path.Join(location, name+strconv.FormatInt(etime.Now().RoundTo(interval), 10)+".log"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0700)
			timer.Reset(time.Duration(interval) * time.Second)
		}
	}
}
