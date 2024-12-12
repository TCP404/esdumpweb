package initial

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const logFileName = "esdumpweb.log"

var (
	logDir            string
	logManagerIns     *LoggerManager
	logManagerInsOnce sync.Once
)

func makeLogDir() {
	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Fatal(err.Error())
	}
	logDir = filepath.Join(userCacheDir, "esdumpweb")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatal(err.Error())
	}
}

type LoggerManager struct {
	logFile     *os.File
	logFileOnce sync.Once
	logFilePath string
	logTime     time.Time
}

func WireLoggerManager() *LoggerManager {
	logManagerInsOnce.Do(func() {
		logManagerIns = &LoggerManager{}
		logManagerIns.init()
	})
	return logManagerIns
}

func (l *LoggerManager) init() {
	l.logFileOnce.Do(func() {
		var err error
		l.logTime = time.Now()
		l.logFilePath = filepath.Join(logDir, logFileName)
		l.logFile, err = os.OpenFile(l.logFilePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("failed to initilize logger. %v", err.Error()))
		}
		logger := slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stdout, l.logFile), &slog.HandlerOptions{Level: slog.LevelDebug}))
		// logger := slog.New(slog.NewTextHandler(io.MultiWriter(os.Stdout, l.logFile), &slog.HandlerOptions{Level: slog.LevelDebug}))
		slog.SetDefault(logger)
	})
}

func (l *LoggerManager) GetLogFilePath() string {
	return l.logFilePath
}

func (l *LoggerManager) GetLogTime() (t time.Time) {
	return l.logTime
}

func (l *LoggerManager) GetFileContent() ([]byte, error) {
	return io.ReadAll(l.logFile)
}

func (l *LoggerManager) Read(p []byte) (n int, err error) {
	return l.logFile.Read(p)
}

func (l *LoggerManager) Write(p []byte) (n int, err error) {
	return l.logFile.Write(p)
}

func (l *LoggerManager) Close() {
	l.logFile.Close()
}
