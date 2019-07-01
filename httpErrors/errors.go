package errors

import "strconv"

type HTTPResponse struct {
	Response interface{}
	Code     int
}

type HTTPError struct {
	Err string
	*HTTPResponse
}

func New(s string, r ...HTTPResponse) error {
	if len(r) == 0 {
		return HTTPError{Err: s}
	}

	return HTTPError{Err: s, HTTPResponse: &r[0]}
}

func (h HTTPError) Error() string {
	if h.HTTPResponse == nil {
		return h.Err
	}

	return "[" + strconv.Itoa(h.Code) + "]: " + h.Err
}
