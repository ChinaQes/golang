// 2020-11-05
// 作者: qes
// -v: 0.1

package qespath

import (
	"os"
	"path/filepath"
	"strings"
)

//fileType *
type fileType uint8

const (
	file fileType = 0
	dir  fileType = 1
	all  fileType = 2
)

//getAllPath *
func getAllPath(path string, fileType fileType) (dirList []string, err error) {

	err = filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			switch fileType {
			case file:
				if !info.IsDir() {
					dirList = append(dirList, path)
					return nil
				}
			case dir:
				if info.IsDir() {
					dirList = append(dirList, path)
					return nil
				}
			case all:
				dirList = append(dirList, path)
				return nil
			}
			return nil
		})

	return dirList, err
}

//GetAllDirPath *
func GetAllDirPath(path string) (dirList []string, err error) {
	return getAllPath(path, dir)
}

//GetAllFilePath *
func GetAllFilePath(path string) (dirList []string, err error) {
	return getAllPath(path, file)
}

//GetAllPath *
func GetAllPath(path string) (dirList []string, err error) {
	return getAllPath(path, all)
}

//GetAllEnv *
func GetAllEnv(variableName string) []string {
	env := os.Getenv(variableName)
	lis := strings.Split(env, ";")
	return lis
}
