package deploy

import "github.com/nomos/promise"

type IDeployProcedure interface {
	UnmarshalFrom()
	MarshalTo()
	Process()promise.Promise
}

