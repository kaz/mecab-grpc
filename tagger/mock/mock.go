package mock

type (
	MockTagger struct{}
)

func New() *MockTagger {
	return &MockTagger{}
}

func (t *MockTagger) Parse(input string) (string, error) {
	return "This is mock implementation. You must select tagger implementation with `-tags ***` on build.\n", nil
}

func (t *MockTagger) Close() error {
	return nil
}
