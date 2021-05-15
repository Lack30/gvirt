package spec

type Document interface {
	Unmarshal(doc string) error
	Marshal() (string, error)
}
