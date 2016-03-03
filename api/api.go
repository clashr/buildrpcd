package api

type Args struct {
	Language, Dialect, Contents string
}

type Build struct {
}

type Result struct {
	CCOut, Binary []byte
}

func (b *Build) Compile(args Args, result *Result) error {
	return compile(args, result)
}
