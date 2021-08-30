package util

import (
	"fmt"
	"os"
	"path/filepath"
)

// FileExists Check if File / Directory Exists
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
// FindExecDir Find Directory where the program starts
func FindExecDir() (string, error) {
	//找到启动文件路径
	execFile,err := filepath.Abs(os.Args[0])

	if err != nil {
		return  "",fmt.Errorf("can't get absolute path for executable: %v",err)
	}

	path := filepath.Dir(execFile)

	realFile,err := filepath.EvalSymlinks(execFile)

	if err != nil{
		//windows
		if _,err = os.Stat(execFile +".exe"); err == nil {
			realFile = filepath.Clean(execFile+".exe")
		}
	}

	if err == nil &&  realFile != execFile {
		path = filepath.Dir(realFile)
	}
	return  path,nil
}
