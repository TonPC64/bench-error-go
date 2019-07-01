package errors

import "strconv"

type HttpResponse struct {
	Response interface{}
	Code     int
}

type HTTPError struct {
	Err string
	*HttpResponse
}

func New(s string, r ...HttpResponse) error {
	if len(r) == 0 {
		return HTTPError{Err: s}
	}

	return HTTPError{Err: s, HttpResponse: &r[0]}
}

func (h HTTPError) Error() string {
	if h.HttpResponse == nil {
		return h.Err
	}

	return "[" + strconv.Itoa(h.Code) + "]: " + h.Err
}
