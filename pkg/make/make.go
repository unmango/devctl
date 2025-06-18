package make

type SetupResult = int

const (
	SetupNoRebuild SetupResult = iota - 1
	SetupFailure
	SetupSuccess
)
