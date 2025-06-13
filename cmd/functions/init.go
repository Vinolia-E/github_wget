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
	ConvertLinks bool
	// ConvertedLinks []string
	RejectFlag string
	ExcludeFlag string
	Reject bool
	// Rejected []string
	// Excluded []string
	Exclude bool
}
