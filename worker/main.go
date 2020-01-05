package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	_ "github.com/mattn/go-sqlite3"

	ydqueue "github.com/satom9to5/youtube-dl-queue"
)

var (
	sqlitePath    = flag.String("sqlite_path", "", "sqlite path")
	pidfilePath   = flag.String("pidfile_path", "", "pidfile path")
	youtubeDlPath = flag.String("youtubedl_path", "", "Youtube-DL path")
	ffmpegPath    = flag.String("ffmpeg_path", "", "ffmpeg path")
	logDirectory  = flag.String("log_directory", "", "output log directory")

	debug = flag.Bool("debug", false, "debug mode")
)

func main() {
	flag.Parse()
	checkFlags()

	// Wait Signal
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGQUIT, syscall.SIGTERM)

	defer shutdown()

	if err := dispatcher(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	// Server shutdown
	s := <-ch
	fmt.Printf("signal catched: %s\n", s)
}

func checkFlags() {
	emptyFlags := []string{}
	if *sqlitePath == "" {
		emptyFlags = append(emptyFlags, "sqlite_path")
	}
	if *pidfilePath == "" {
		emptyFlags = append(emptyFlags, "pidfile_path")
	}
	if *youtubeDlPath == "" {
		emptyFlags = append(emptyFlags, "youtubedl_path")
	}
	if *ffmpegPath == "" {
		emptyFlags = append(emptyFlags, "ffmpeg_path")
	}
	if *logDirectory == "" {
		emptyFlags = append(emptyFlags, "log_directory")
	}
	if len(emptyFlags) > 0 {
		shutdown(errors.New(strings.Join(emptyFlags, "/") + "is empty."))
	}
}

func dispatcher() error {
	db, err := sql.Open("sqlite3", *sqlitePath)
	if err != nil {
		return err
	}

	ydqueue.SetLogDirectory(*logDirectory)

	_, err = ydqueue.Start(
		db,
		*pidfilePath,
		*youtubeDlPath,
		*ffmpegPath,
	)

	return err
}

func shutdown(v ...interface{}) {
	ydqueue.Stop()

	exitStatus := 0

	if len(v) > 0 {
		exitStatus = 1
		fmt.Println(v...)
	}

	fmt.Println("Shutdown queue process...")

	os.Exit(exitStatus)
}
