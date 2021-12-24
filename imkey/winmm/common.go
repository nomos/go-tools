package winmm

import "golang.org/x/sys/windows"

const (
	WinMMDllName = "winmm.dll"
)

type WinMMDLL struct {
	WinMM           *windows.DLL
	timeBeginPeriod *windows.Proc
	timeEndPeriod   *windows.Proc
}
func LoadWinMMDll() (*WinMMDLL, error) {
	temp := windows.LazyDLL{
		Name:   WinMMDllName,
		System: false,
	}
	WinMM := &windows.DLL{
		Name:   temp.Name,
		Handle: windows.Handle(temp.Handle()),
	}
	timeBeginPeriod, err := WinMM.FindProc("timeBeginPeriod")
	if err != nil {
		return nil, err
	}
	timeEndPeriod, err := WinMM.FindProc("timeEndPeriod")
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	ret := &WinMMDLL{
		WinMM:            WinMM,
		timeBeginPeriod:   timeBeginPeriod,
		timeEndPeriod: timeEndPeriod,
	}
	return ret, nil

}

func (this *WinMMDLL) Release() error {
	return this.WinMM.Release()
}

func (this *WinMMDLL) TimeBeginPeriod(millie uintptr){
	this.timeBeginPeriod.Call(millie)
}

func (this *WinMMDLL) TimeEndPeriod(millie uintptr){
	this.timeEndPeriod.Call(millie)
}