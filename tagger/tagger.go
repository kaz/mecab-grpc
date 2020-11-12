package tagger

type (
	Tagger interface {
		Parse(string) (string, error)
	}
)
