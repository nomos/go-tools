package kernel32

import (
	"golang.org/x/sys/windows"
)

const (
	HIGH_PRIORITY_CLASS uintptr = 0x00000080
	Kernel32DllName             = "kernel32.dll"
)

func LoadKernel32Dll() (*Kernel32DLL, error) {
	temp := windows.LazyDLL{
		Name:   Kernel32DllName,
		System: false,
	}
	Kernel32 := &windows.DLL{
		Name:   temp.Name,
		Handle: windows.Handle(temp.Handle()),
	}
	getCurrentProcess, err := Kernel32.FindProc("GetCurrentProcess")
	if err != nil {
		return nil, err
	}
	getCurrentProcessId, err := Kernel32.FindProc("GetCurrentProcessId")
	if err != nil {
		return nil, err
	}
	setPriorityClass, err := Kernel32.FindProc("SetPriorityClass")
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	ret := &Kernel32DLL{
		Kernel32:            Kernel32,
		getCurrentProcess:   getCurrentProcess,
		getCurrentProcessId: getCurrentProcessId,
		setPriorityClass:    setPriorityClass,
	}
	return ret, nil

}

type Kernel32DLL struct {
	Kernel32            *windows.DLL
	context             uintptr
	getCurrentProcess   *windows.Proc
	getCurrentProcessId *windows.Proc
	setPriorityClass    *windows.Proc
}

func (this *Kernel32DLL) GetCurrentProcess() uintptr {
	ret, _, _ := this.getCurrentProcess.Call()
	return ret
}

func (this *Kernel32DLL) GetCurrentProcessId() uintptr {
	ret, _, _ := this.getCurrentProcessId.Call()
	return ret
}

func (this *Kernel32DLL) SetPriorityClass(process uintptr, prior uintptr) bool {
	this.setPriorityClass.Call(process, prior)
	return true
}
func (this *Kernel32DLL) Release() error {
	return this.Kernel32.Release()
}
