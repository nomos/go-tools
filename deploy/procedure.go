package deploy

import "github.com/nomos/go-lokas/promise"

type IDeployProcedure interface {
	UnmarshalFrom()
	MarshalTo()
	Process()promise.Promise
}

