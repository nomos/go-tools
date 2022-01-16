package imkey

type ITask interface {
	Sleep(duration int64)bool
	Run()
	TaskOff()
	TaskOn()
}