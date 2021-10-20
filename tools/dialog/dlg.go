package dialog

import (
	"errors"
	"fmt"
)

var ErrCancelled = errors.New("Cancelled")

var Cancelled = ErrCancelled

type Dlg struct {
	Title string
}

type MsgBuilder struct {
	Dlg
	Msg string
}

func Message(format string, args ...interface{}) *MsgBuilder {
	return &MsgBuilder{Msg: fmt.Sprintf(format, args...)}
}

func (b *MsgBuilder) Title(title string) *MsgBuilder {
	b.Dlg.Title = title
	return b
}

func (b *MsgBuilder) YesNo() bool {
	return b.yesNo()
}

func (b *MsgBuilder) Info() {
	b.info()
}

func (b *MsgBuilder) Error() {
	b.error()
}

type FileFilter struct {
	Desc       string
	Extensions []string
}

type FileBuilder struct {
	Dlg
	StartDir string
	Filters  []FileFilter
}

func File() *FileBuilder {
	return &FileBuilder{}
}

func (b *FileBuilder) Title(title string) *FileBuilder {
	b.Dlg.Title = title
	return b
}

func (b *FileBuilder) Filter(desc string, extensions ...string) *FileBuilder {
	filt := FileFilter{desc, extensions}
	if len(filt.Extensions) == 0 {
		filt.Extensions = append(filt.Extensions, "*")
	}
	b.Filters = append(b.Filters, filt)
	return b
}

func (b *FileBuilder) SetStartDir(startDir string) *FileBuilder {
	b.StartDir = startDir
	return b
}

func (b *FileBuilder) Load() (string, error) {
	return b.load()
}

func (b *FileBuilder) Save() (string, error) {
	return b.save()
}

type DirectoryBuilder struct {
	Dlg
	StartDir string
}

func Directory() *DirectoryBuilder {
	return &DirectoryBuilder{}
}

func (b *DirectoryBuilder) Browse() (string, error) {
	return b.browse()
}

func (b *DirectoryBuilder) Title(title string) *DirectoryBuilder {
	b.Dlg.Title = title
	return b
}

func (b *DirectoryBuilder) SetStartDir(dir string) *DirectoryBuilder {
	b.StartDir = dir
	return b
}
