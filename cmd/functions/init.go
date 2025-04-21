package functions

type ArgValues struct {
	URL []string
	Flag string
	OutputFile string
	Path string
	Background bool
	InputFile string
	IsMirror bool
	RateLimit string
	HasFlag bool
}

type MirrorValues struct {
	Convert bool
	Exclude bool
	Reject bool
	Rejected []string
	Excluded []string
	ConvertPath string
}