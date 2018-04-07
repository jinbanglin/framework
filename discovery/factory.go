package discovery

import (
	"io"

	"github.com/jinbanglin/moss"
)

// Factory is a function that converts an instance string (e.g. host:port) to a
// specific endpoint. Instances that provide multiple endpoints require multiple
// factories. A factory also returns an io.Closer that's invoked when the
// instance goes away and needs to be cleaned up. Factories may return nil
// closers.
//
// Users are expected to provide their own factory functions that assume
// specific transports, or can deduce transports by parsing the instance string.
type Factory func(instance string) (moss.Endpoint, io.Closer, error)