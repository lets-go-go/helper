package rerun

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/lets-go-go/helper/file"
	"github.com/lets-go-go/logger"
)

func init() {
	cuurentPid := fmt.Sprint(os.Getpid())
	fileHash, err := file.GetFileHash(os.Args[0])
	if err != nil {
		return
	}

	pidFilename := fmt.Sprintf("%s.pid", fileHash.MD5)

	pidPath := filepath.Join(os.TempDir(), pidFilename)

	if ret, pid := CheckProcExsit(pidPath); ret {
		logger.Warnf("process has running.pid=%v", pid)
		os.Exit(1)
	} else {
		pidFile, _ := os.Create(pidPath)
		defer pidFile.Close()
		pidFile.WriteString(cuurentPid)
		logger.Infof("current process id=%v", cuurentPid)
	}
}

// CheckProcExsit Check Process Exsit
func CheckProcExsit(pidPath string) (bool, string) {
	pidFile, err := os.Open(pidPath)
	defer pidFile.Close()
	pidStr := ""
	if err == nil {
		filePid, err := ioutil.ReadAll(pidFile)
		if err == nil {
			pidStr = fmt.Sprintf("%s", filePid)
			pid, _ := strconv.Atoi(pidStr)
			_, err := os.FindProcess(pid)
			if err == nil {
				return true, pidStr
			}
		}
	}
	return false, pidStr
}
