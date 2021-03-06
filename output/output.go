/**
 * @Author: guobob
 * @Description:
 * @File:  output.go
 * @Version: 1.0.0
 * @Date: 2022/3/27 22:06
 */

package output

import (
	"fmt"
)

type Output interface {
	WriteData(dbName, tableName string, buff []byte) error
	Sync() error
	Close()
}

type TableFiles struct {
	files       map[string]*WriteFile
	filePath    string
	filePrefix  string
	maxFileSize uint64
	maxFileNum  uint64
	sync        bool
}

func NewTableFiles(sync bool, maxFileSize, maxFileNum uint64, path, filePrefix string) *TableFiles {
	/*
		if len(strings.TrimSpace(filePrefix)) == 0 {
			ts := time.Now()
			filePrefix = fmt.Sprintf("%v%02d%02d", ts.Year(), ts.Month(), ts.Day())
		}
	*/
	return &TableFiles{
		sync:        sync,
		maxFileSize: maxFileSize,
		maxFileNum:  maxFileNum,
		filePath:    path,
		filePrefix:  filePrefix,
		files:       make(map[string]*WriteFile),
	}
}

func (tf *TableFiles) Close() {
	for _, v := range tf.files {
		v.close()
	}
}

func (tf *TableFiles) Sync() error {
	var err error
	if tf.sync {
		return nil
	}
	for _, v := range tf.files {
		err = v.fp.Sync()
		if err != nil {
			return err
		}
	}
	return err
}

func (tf *TableFiles) WriteData(dbName, tableName string, buff []byte, startfileNo uint64) error {
	var err error
	key := fmt.Sprintf("%v.%v", dbName, tableName)
	v, ok := tf.files[key]
	if !ok {
		wf := newWriteFile(tf, tableName, dbName, startfileNo)
		tf.files[key] = wf
		v = wf
		v.getFileNo()
		v.generateFileName()
		err = v.openFile()
		if err != nil {
			return err
		}
		v.pos = 0
	}
	/*
		v.buff = buff
		err = v.write()
		if err != nil {
			return err
		}
	*/
	v.WriteAsync(buff)
	if v.checkIfNeedChangeFile() {
		v.close()
		v.getFileNo()
		v.generateFileName()
		err := v.openFile()
		if err != nil {
			return err
		}
		v.pos = 0
		v.currentNum = 0
	}
	return err
}
