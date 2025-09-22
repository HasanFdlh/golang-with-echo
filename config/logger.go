package config

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() {
	// bikin folder logs kalau belum ada
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}

	// nama file log berdasarkan tanggal
	logFile := filepath.Join("logs", time.Now().Format("2006-01-02")+".log")

	// lumberjack rotate otomatis per hari (maxAge=7 hari)
	logger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    50, // MB
		MaxBackups: 30,
		MaxAge:     14,   // hari simpan log
		Compress:   true, // kompres log lama (gzip)
	}

	log.SetOutput(logger)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println("[SYSTEM] Logger initialized, writing to", logFile)
}
