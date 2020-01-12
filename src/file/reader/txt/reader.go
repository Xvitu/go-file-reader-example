package txt

import (
	"os"
	"strings"
)

type File struct {
	path    string
	content []byte
}

type txtReader interface {
	Read() (ok bool, err error)
	AsStringSlice() (content []string, err error)
	asString() (content string)
}

func New(path string) *File {
	return &File{path: path}
}

func (f *File) Read() (ok bool, err error) {
	ok = true
	file, err := os.Open(f.path)

	defer file.Close()

	if err != nil {
		ok = false
		return
	}

	byteContent := make([]byte, 5000)
	_, err = file.Read(byteContent)

	f.content = byteContent

	if err != nil {
		ok = false
		return
	}

	return
}

func (f File) asString() (content string) {
	content = string(f.content)

	return
}

func (f File) AsStringSlice() (content []string) {
	textFromFile := f.asString()

	content = strings.Split(textFromFile, " ")
	return
}
