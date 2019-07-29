package fission_io

type (
	Object interface {
		Validate() error
		Merge(new Object, old Object) error
	}
)
