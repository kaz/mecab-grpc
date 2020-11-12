package tagger

type (
	Config map[string]string

	Tagger interface {
		Parse(string) (string, error)
		Close() error
	}
)
