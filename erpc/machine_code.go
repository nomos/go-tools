package erpc

import (
	"github.com/super-l/machine-code/machine"
)

func (this *App) GetMachineData()  {
	this.MachineData=machine.GetMachineData()
}

func (this *App) MacAddr()string {
	return this.MachineData.Mac
}

func (this *App) CpuId()string {
	return this.MachineData.CpuId
}

func (this *App) PlatformUUID()string {
	return this.MachineData.PlatformUUID
}

func (this *App) SerialNumber()string {
	return this.MachineData.SerialNumber
}