package middleware

import (
	"io"
)

type Router interface {
	Route([]byte)
}

type Parser interface {
	Parse(io.ReadCloser, Router)
}