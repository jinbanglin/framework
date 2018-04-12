package http

import (
	"context"
	"net/http"
)

type DecodeRequestFunc func(context.Context, *http.Request) (request interface{}, err error)

type EncodeResponseFunc func(context.Context, http.ResponseWriter, interface{}) error
