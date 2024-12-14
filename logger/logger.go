package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
	"tool/unit"
)

const (
	DAY     = time.Hour * 24
	WEEK    = DAY * 7
	MONTH   = DAY * 30
	_TRACE_ = "[TAC]"
	_DEBUG_ = "[DBG]"
	_INFO_  = "[INF]"
	_WARN_  = "[WRN]"
	_ERROR_ = "[ERR]"
	_PANIC_ = "[PNC]"
	_SEP_   = "."
)

func Trace(v ...any) { g_loggge.Trace(v...) }
func Debug(v ...any) { g_loggge.Debug(v...) }
func Info(v ...any)  { g_loggge.Info(v...) }
func Warn(v ...any)  { g_loggge.Warn(v...) }
func Error(v ...any) { g_loggge.Error(v...) }
func Panic(v ...any) { g_loggge.Panic(v...) }

func (r *Logger) Trace(v ...any) {
	r.worker.Push(&LogOutPut{l: r, node: &logNode{data: v, prefix: _TRACE_, pushTime: time.Now()}})
}
func (r *Logger) Debug(v ...any) {
	r.worker.Push(&LogOutPut{l: r, node: &logNode{data: v, prefix: _DEBUG_, pushTime: time.Now()}})
}
func (r *Logger) Info(v ...any) {
	r.worker.Push(&LogOutPut{l: r, node: &logNode{data: v, prefix: _INFO_, pushTime: time.Now()}})
}
func (r *Logger) Warn(v ...any) {
	r.worker.Push(&LogOutPut{l: r, node: &logNode{data: v, prefix: _WARN_, pushTime: time.Now()}})
}
func (r *Logger) Error(v ...any) {
	r.worker.Push(&LogOutPut{l: r, node: &logNode{data: v, prefix: _ERROR_, pushTime: time.Now()}})
}
func (r *Logger) Panic(v ...any) {
	r.worker.Push(&LogOutPut{l: r, node: &logNode{data: v, prefix: _PANIC_, pushTime: time.Now()}})
}

type (
	fileEntry struct {
		createTime time.Time
		path       string
	}
	logNode struct {
		data     []any
		prefix   string
		pushTime time.Time
	}
	Logger struct {
		out         *os.File
		worker      *unit.WorkerPool
		buf         []byte
		createdFile []*fileEntry
	}
)

func CloseLog() {
	g_loggge.Close()
}

func (r *Logger) Close() {
	r.worker.Close()
	if nil != r.out {
		r.out.Sync()
	}
}

var g_loggge *Logger

func NewLogger(intervalTime, survivalTime time.Duration) (err error) {
	g_loggge, err = NewSyncLog(filepath.Join(filepath.Dir(os.Args[0]), "logs"), intervalTime, survivalTime)
	return err
}
func NewSyncLog(dir string, intervalTime, survivalTime time.Duration) (*Logger, error) {
	if info, err := os.Stat(dir); nil == info || os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0755); nil != err {
			return nil, err
		}
	}

	if intervalTime < time.Minute {
		intervalTime = time.Minute
	}

	for survivalTime <= intervalTime {
		survivalTime += time.Minute * 5
	}

	r := &Logger{worker: unit.NewWorkerPool(-1, 1)}
	r.worker.SetLogger(r)

	fileLink := filepath.Join(dir, "log")
	if matches, err := filepath.Glob(fileLink + _SEP_ + "*"); nil == err {
		for _, match := range matches {
			info, err := os.Stat(match)
			if nil == err && nil != info {
				r.createdFile = append(r.createdFile, &fileEntry{createTime: info.ModTime(), path: match})
			}
		}
		sort.Slice(r.createdFile, func(i, j int) bool { return r.createdFile[i].createTime.Before(r.createdFile[j].createTime) })
	}

	if errs := r.rotateFile(fileLink, survivalTime); len(errs) > 0 {
		return nil, errs[0]
	}

	timer := time.NewTimer(time.Now().Truncate(intervalTime).Add(intervalTime).Sub(time.Now()))
	go func() {
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				r.worker.Push(&RotateFile{
					l:            r,
					fileLink:     fileLink,
					survivalTime: survivalTime,
				})
				timer.Reset(time.Now().Truncate(intervalTime).Add(intervalTime).Sub(time.Now()))
			}
		}
	}()
	return r, nil
}

type RotateFile struct {
	l            *Logger
	fileLink     string
	survivalTime time.Duration
}

func (r *RotateFile) Callback() {
	r.l.rotateFile(r.fileLink, r.survivalTime)
}

func (r *Logger) rotateFile(fileLink string, survivalTime time.Duration) (ret []error) {
	now := time.Now()
	newFile := fileLink + _SEP_ + strings.ReplaceAll(now.Format(time.RFC3339)[:len("2006-01-02T15:04:05")], ":", "-")
	if info, err := os.Stat(newFile); nil == info || os.IsNotExist(err) {
		r.createdFile = append(r.createdFile, &fileEntry{createTime: now, path: newFile})
	}
	fp, err := os.OpenFile(newFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if nil != err {
		ret = append(ret, err)
		return
	}
	os.Remove(fileLink)
	if err = os.Link(newFile, fileLink); nil != err {
		ret = append(ret, err)
	}

	if nil != r.out {
		if err = r.out.Close(); nil != err {
			ret = append(ret, err)
		}
	}
	r.out = fp

	delNum := 0
	for i := 0; i < len(r.createdFile); i++ {
		if r.createdFile[i].createTime.Add(survivalTime).After(now) {
			break
		}
		if err = os.Remove(r.createdFile[i].path); nil == err {
			delNum++
		} else {
			ret = append(ret, err)
		}
	}
	r.createdFile = r.createdFile[delNum:]
	return
}

func itoa(buf *[]byte, i int, wid int) {
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}

func formatPrefix(buf *[]byte, t time.Time, prefix string) {
	h, m, s := t.Clock()
	itoa(buf, h, 2)
	*buf = append(*buf, ':')
	itoa(buf, m, 2)
	*buf = append(*buf, ':')
	itoa(buf, s, 2)
	*buf = append(*buf, '.')
	itoa(buf, t.Nanosecond()/1e3, 6)
	*buf = append(*buf, ' ')
	*buf = append(*buf, prefix...)
	*buf = append(*buf, ' ')
}

type LogOutPut struct {
	l    *Logger
	node *logNode
}

func (r *LogOutPut) Callback() {
	r.l.output(r.node)
}
func (l *Logger) output(node *logNode) {
	l.buf = l.buf[:0]
	formatPrefix(&l.buf, node.pushTime, node.prefix)
	l.buf = fmt.Appendln(l.buf, node.data...)
	if len(l.buf) > 1 && l.buf[len(l.buf)-2] == '\n' {
		l.buf = l.buf[:len(l.buf)-1]
	}
	l.out.Write(l.buf)
}
