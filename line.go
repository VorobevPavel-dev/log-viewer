package main

import "os"

type Line struct {
	Start int
	End   int
}

// IsIn checks if provided position inside current Line
func (l *Line) IsIn(position int) bool {
	return position <= l.End && position >= l.Start
}

// Down checks if provided position is somewhere in next lines
func (l *Line) Down(position int) bool {
	return position > l.End
}

// Up checks if provieded position is somewhere in previous lines
func (l *Line) Up(position int) bool {
	return position < l.Start
}

func (l *Line) ReadFromFile(file *os.File) []byte {
	var result = make([]byte, l.End-l.Start+1)
	file.Seek(int64(l.Start), 0)
	file.Read(result)
	return result
}
