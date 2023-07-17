package errors

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

// Define Error
var (
	ErrMarshal   = status.New(0, "Marshal failed").Err()
	ErrUnmarshal = status.New(1, "Unmarshal failed").Err()
)

// ErrorWithMessage wraps detail error
func ErrorWithMessage(err error, message string) error {
	s, ok := status.FromError(err)
	if !ok {
		return errors.Wrap(err, message)
	}

	grpcStatus := status.New(s.Code(), s.Message()+" "+message)
	return grpcStatus.Err()
}
