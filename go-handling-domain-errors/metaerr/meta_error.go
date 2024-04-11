package metaerr

import "errors"

// errMetaData is a struct that contains an error and metadata.
type errMetaData struct {
	// err is the error that occurred.
	err error
	// metadata is the metadata associated with the error.
	metadata []any
}

// return original error message
// implement the error interface
func (e *errMetaData) Error() string {
	return e.err.Error()
}

// allow unwrapping of the error compatible with Go error wrapping mechanism
func (e *errMetaData) Unwrap() error {
	return e.err
}

func WihMetadata(err error, pairs ...any) error {
	return &errMetaData{
		err:      err,
		metadata: pairs,
	}
}

func GetMetadata(err error) []any {
	data := []any{}

	for err != nil {
		if e, ok := err.(*errMetaData); ok {
			data = append(data, e.metadata...)
		}

		err = errors.Unwrap(err)
	}

	return data
}
