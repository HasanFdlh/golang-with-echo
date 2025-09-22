package middleware

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// bikin folder logs kalau belum ada
		if _, err := os.Stat("logs"); os.IsNotExist(err) {
			os.Mkdir("logs", 0755)
		}

		// nama file log = tanggal hari ini
		logFile := filepath.Join("logs", time.Now().Format("2006-01-02")+".log")

		// buka file append
		f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Printf("[ERROR] Gagal buka file log: %v", err)
			next.ServeHTTP(w, r)
			return
		}
		defer f.Close()

		// bikin logger khusus file ini
		logger := log.New(f, "", log.LstdFlags)

		start := time.Now()
		next.ServeHTTP(w, r) // proses ke handler

		// catat log request
		logger.Printf("%s %s %s %s",
			r.RemoteAddr,
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}
