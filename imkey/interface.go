package imkey

import "time"

type ITask interface {
	Sleep(duration time.Duration)
	Run()
	TaskOff()
}