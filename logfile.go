package main

import (
	"fmt"
	"os"
	"path"
)

type Logfile struct {
	Filepath      string `json:"filepath"`
	Filename      string `json:"filename"`
	lines         *[]Line
	currentOffset int64
	file          *os.File
}

func InitLogfile(filepath string) (*Logfile, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("file %s cannot be openned", filepath)
	}
	result := Logfile{
		Filepath:      filepath,
		Filename:      path.Base(filepath),
		currentOffset: 0,
		lines:         determineLines(file, 0),
		file:          file,
	}
	return &result, nil
}

func (l *Logfile) Getline(linenum int) ([]byte, error) {
	if linenum > len(*l.lines)-1 {
		return nil, fmt.Errorf("file contains only %d lines. Requested line: %d", len(*l.lines)-1, linenum)
	}
	return (*l.lines)[linenum].ReadFromFile(l.file), nil
}

func (l *Logfile) Close() {
	l.file.Close()
}

func determineLines(file *os.File, startPosition int) *[]Line {
	var (
		stepSize       = 100
		batch          = make([]byte, stepSize)
		globalPosition = 0
		tempLine       = Line{
			Start: 0,
		}
		fileEndPos = 0
		result     = &[]Line{}
	)
	file.Seek(int64(startPosition), 0)
	readBytes, err := file.Read(batch)
	for err == nil {
		if readBytes < stepSize {
			fileEndPos = globalPosition + readBytes
		}
		for i := 0; i < readBytes; i++ {
			if batch[i] == '\n' {
				tempLine.End = globalPosition + i - 1
				*result = append((*result), tempLine)
				tempLine = Line{
					Start: globalPosition + i + 1,
				}
			}
		}
		globalPosition += stepSize
		readBytes, err = file.Read(batch)
	}
	tempLine.End = fileEndPos
	*result = append(*result, tempLine)
	return result
}
