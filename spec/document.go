package spec

type Document interface {
	UnmarshalX(doc string) error
	MarshalX() (string, error)
}
