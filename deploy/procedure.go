package deploy

import "github.com/nomos/go-promise"

type IDeployProcedure interface {
	UnmarshalFrom()
	MarshalTo()
	Process()promise.Promise
}

