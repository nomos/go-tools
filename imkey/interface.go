package imkey

type ITask interface {
	Sleep(duration int)bool
	Run()
	TaskOff()
}