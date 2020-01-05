package action

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/satom9to5/pidfile"
)

type worker struct {
	SqlitePath    string `mapstructure:"sqlite_path"`
	PidfilePath   string `mapstructure:"pidfile_path"`
	YoutubeDlPath string `mapstructure:"youtubedl_path"`
	FFMpegPath    string `mapstructure:"ffmpeg_path"`
	LogDirectory  string `mapstructure:"log_directory"`
	Browser       string `mapstructure:"browser"`
	Dev           bool   `mapstructure:"dev"`
}

func (w worker) String() string {
	return fmt.Sprintf(
		"SqlitePath: \"%s\", PidfilePath: \"%s\", YoutubeDlPath: \"%s\", FFMpegPath: \"%s\", LogDirectory: \"%s\", Browser: \"%s\"",
		w.SqlitePath,
		w.PidfilePath,
		w.YoutubeDlPath,
		w.FFMpegPath,
		w.LogDirectory,
		w.Browser,
	)
}

func (w *worker) Start() (pid int, err error) {
	pidfile.Initialize(w.PidfilePath)

	if res, _ := w.CheckRunning(); res {
		return 0, errors.New("Worker already started.")
	}

	cmd, err := w.command()
	if err != nil {
		return 0, errors.New("Command Parameter error.")
	}

	w.setLog(cmd)
	w.setFlag(cmd)

	err = cmd.Start()

	// wait start child process
	for i := 0; i < 10; i++ {
		if proc, _ := pidfile.GetProcess(); proc != nil {
			return proc.Pid, nil
		}
		time.Sleep(time.Second)
	}

	return 0, errors.New("Falied Worker start.")
}

func (w *worker) Stop() error {
	pidfile.Initialize(w.PidfilePath)

	proc, err := pidfile.GetProcess()
	if err != nil {
		return err
	}

	switch runtime.GOOS {
	case "windows":
		//  cannot use Signal when Windows, remove pidfile here.
		pidfile.Remove()
		return proc.Kill()
	default:
		return proc.Signal(os.Interrupt)
	}
}

func (w *worker) CheckRunning() (bool, error) {
	pidfile.Initialize(w.PidfilePath)

	process, err := pidfile.GetProcess()
	if err != nil {
		return false, err
	}

	return process != nil, nil
}

func (w *worker) setLog(cmd *exec.Cmd) error {
	sep := string(os.PathSeparator)

	if logFile, err := w.createLogFile(w.LogDirectory + sep + "worker_stdout.log"); err == nil {
		cmd.Stdout = logFile
	} else {
		return err
	}

	if logFile, err := w.createLogFile(w.LogDirectory + sep + "worker_stderr.log"); err == nil {
		cmd.Stderr = logFile
	} else {
		return err
	}

	return nil
}

func (w *worker) createLogFile(path string) (io.Writer, error) {
	return os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
}

func (w *worker) command() (*exec.Cmd, error) {
	sep := string(os.PathSeparator)

	commonParams := []string{
		"-sqlite_path", w.SqlitePath,
		"-pidfile_path", w.PidfilePath,
		"-youtubedl_path", w.YoutubeDlPath,
		"-ffmpeg_path", w.FFMpegPath,
		"-log_directory", w.LogDirectory,
	}

	if w.Dev {
		mainGoPath := strings.Join([]string{
			"../worker", "main.go",
		}, sep)

		params := append([]string{
			"run", mainGoPath, "-debug",
		}, commonParams...)

		return exec.Command("go", params...), nil
	} else {
		return exec.Command(w.workerPath(), commonParams...), nil
	}
}
