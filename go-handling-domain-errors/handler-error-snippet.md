```
// Definition
type GenericError struct {
    Code string
    Message string
}

func (e *GenericError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Definition specific error
type ValidationError struct {
    GenericError
}

func NewValidationError(code, msg string) *ValidationError {
    return &ValidationError{GenericError{
        Code: code,
        Message: msg,
    }}
}

// Implementation of specific error
// ErrSomethingIsEmpty is a representation of the specific domain error
// and can be defined on a package level or in some function as well.
var ErrSomethingIsEmpty = NewValidationError("something_is_empty",
    "something should have a specific value and cannot be empty")

func (e *DomainEntity) DoSomething(val string) error {
    if val == "" {
        return ErrSomethingIsEmpty
    }

    // do the work …

    return nil
}



// handleRequest is some http request handler in infra layer
func handleRequest(r http.Request, w http.ResponseWriter) {
    err := domainService.DoSomething()
    if err != nil {
        writeError(err, w)
        return
    }
}

// writeError is a function which translates domain errors to status codes
func writeError(err error, w http.ResponseWriter) {
    var errAuth *domain.UnauthenticatedError
    if errors.As(err, &errAuth) {
        writeTypedError(w, http.StatusUnauthorized, errAuth)
        return
    }

    var errNotFound *domain.NotFoundError
    if errors.As(err, &errNotFound) {
        writeTypedError(w, http.StatusNotFound, errNotFound)
        return
    }

    // … continue with other error types

    // when no typed error received, return generic internal error
    // since we don't return any details to a client, it makes sense
    // to log actual error so the error context is persisted.
    writeTypedError(w, http.StatusInternalServerError, domain.ErrInternal)
}

// httpError is an example error JSON representation to return.
type httpError struct {
    Code string `json:"code"`
    Message string `json:"message"`
}

// DomainError is an interface all domain errors should implement.
type DomainError interface {
    Code() string
    Message() string
    Error() string
}

// writeTypedError writes domain error to the response writer.
func writeTypedError(w http.ResponseWriter, code int, domainErr DomainError) {
    errPayload := &httpError{
        Code: domainErr.Code(),
        Message: domainErr.Message(),
    }

    body, _ := json.Marshal(errPayload)
    w.WriteHeader(code)
    w.Write(body)
}
```
