package dialog

import (
	"github.com/nomos/go-lokas/log"
	"os/exec"
	"path"
	"strings"
	"syscall"
)

// osaEscape escapes a string to be used in AppleScript
func osaEscapeString(unescaped string) string {
	escaped := strings.ReplaceAll(unescaped, "\\", "\\\\")
	escaped = strings.ReplaceAll(escaped, "\"", "\\\"")
	escaped = strings.ReplaceAll(escaped, "\n", "\\\n")
	return `"` + escaped + `"`
}

// osaExecute executes AppleScript
func osaExecute(command ...string) (string, error) {
	osa, err := exec.LookPath("osascript")
	if err != nil {
		return "", err
	}

	out, err := exec.Command(osa, "-e", strings.Join(command, "\n")).Output()
	return string(out), err
}

func (b *MsgBuilder) yesNo() bool {
	btn := "Yes"
	//return cocoa.YesNoDlg(b.Msg, b.Dlg.Title)
	out, err := osaExecute(`set T to button returned of (display dialog ` + osaEscapeString(b.Msg) + ` with title ` + osaEscapeString(b.Dlg.Title) + ` buttons {"No", "Yes"} default button ` + osaEscapeString(btn) + `)`)
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return ws.ExitStatus() == 0
		}
	}

	ret := false
	if strings.TrimSpace(out) == "Yes" {
		ret = true
	}

	return ret
}

func osaDialog(title, text, icon string) bool {
	out, err := osaExecute(`display dialog ` + osaEscapeString(text) + ` with title ` + osaEscapeString(title) + ` buttons {"OK"} default button "OK" with icon ` + icon + ``)
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return ws.ExitStatus() == 0
		}
	}

	ret := false
	if strings.TrimSpace(out) == "OK" {
		ret = true
	}

	return ret
}

func (b *MsgBuilder) info() {
	osaDialog(b.Dlg.Title, b.Msg, "note")
}

func (b *MsgBuilder) error() {
	osaDialog(b.Dlg.Title, b.Msg, "stop")
}

//choose file of type {"txt", "jpg", "png"} with prompt ".txt" default location "/Users"
func (b *FileBuilder) load() (string, error) {
	loc := ` default location `
	t := ""
	for _, filt := range b.Filters {
		for _, ext := range filt.Extensions {
			if t==""{
				t = ` of type {`
			}
			if ext == "*" {
			} else {
				ext = strings.TrimLeft(ext,".")
				t += osaEscapeString(ext)
				t += ", "
			}
		}
	}
	t = strings.TrimRight(t,", ")
	if t!=""{
		t += "}"
	}
	log.Infof("ttt",t)
	o, err := osaExecute(`choose file` + t + `with prompt ` + osaEscapeString(b.Dlg.Title) + loc + osaEscapeString(b.StartDir))
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			//ws := exitError.Sys().(syscall.WaitStatus)
			return "", exitError
		}
		return "", ErrCancelled
	}

	out := strings.TrimSpace(o)
	tmp := strings.Split(out, ":")
	outPath := "/" + path.Join(tmp[1:]...)

	return outPath, err
}

//choose file name with prompt ".txt" default location "/Users"
func (b *FileBuilder) save() (string, error) {
	loc := ` default location `
	o, err := osaExecute(`choose file name with prompt ` + osaEscapeString(b.Dlg.Title) + loc + osaEscapeString(b.StartDir))
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			//ws := exitError.Sys().(syscall.WaitStatus)
			return "", exitError
		}
		return "", ErrCancelled
	}

	out := strings.TrimSpace(o)
	tmp := strings.Split(out, ":")
	outPath := "/" + path.Join(tmp[1:]...)

	return outPath, err
}

func (b *FileBuilder) run(save bool) (string, error) {
	//star := false
	//var exts []string
	//for _, filt := range b.Filters {
	//	for _, ext := range filt.Extensions {
	//		if ext == "*" {
	//			star = true
	//		} else {
	//			exts = append(exts, ext)
	//		}
	//	}
	//}
	//if star && save {
	//	/* OSX doesn't allow the user to switch visible file types/extensions. Also
	//	** NSSavePanel's allowsOtherFileTypes property has no effect for an open
	//	** dialog, so if "*" is a possible extension we must always show all files. */
	//	exts = nil
	//}
	//f, err := cocoa.FileDlg(save, b.Dlg.Title, exts, star)
	//if f == "" && err == nil {
	//	return "", ErrCancelled
	//}
	//return f, err
	return "", nil
}

func File111(title, filter string, directory bool) (string, bool, error) {
	f := "file"
	if directory {
		f = "folder"
	}

	t := ""
	if filter != "" {
		t = ` of type {`
		patterns := strings.Split(filter, " ")
		for i, p := range patterns {
			p = strings.Trim(p, "*.")
			t += osaEscapeString(p)
			if i < len(patterns)-1 {
				t += ", "
			}
		}
		t += "}"
	}

	o, err := osaExecute(`choose ` + f + t + ` with prompt ` + osaEscapeString(title))
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return "", ws.ExitStatus() == 0, nil
		}
	}

	out := strings.TrimSpace(o)
	tmp := strings.Split(out, ":")
	outPath := "/" + path.Join(tmp[1:]...)

	return outPath, true, err
}

//choose folder with prompt "指定提示信息" default location "/Users"

func (b *DirectoryBuilder) browse() (string, error) {
	loc := ` default location `
	o, err := osaExecute(`choose folder with prompt ` + osaEscapeString(b.Dlg.Title) + loc + osaEscapeString(b.StartDir))
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			//ws := exitError.Sys().(syscall.WaitStatus)
			return "", exitError
		}
		return "", ErrCancelled
	}

	out := strings.TrimSpace(o)
	tmp := strings.Split(out, ":")
	outPath := "/" + path.Join(tmp[1:]...)

	return outPath, err
}
