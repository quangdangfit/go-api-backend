package errors

import (
	"encoding/json"
	"os"

	"github.com/golang/protobuf/ptypes/any"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/status"
)

// Error constants
const (
	errMsgSep string = ":    "

	// ServiceName : crm service
	ServiceName string = "crm-service"
)

// Error includes Code and Message
type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Service string    `json:"service"`
	s       *spb.Status
}

// Error : implement error.Error()
func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	bs, _ := json.Marshal(e)
	return string(bs)
}

// grpcEncode : encodes to grpc error format
func (e *Error) grpcEncode() {
	grpcStatus := new(spb.Status)
	grpcStatus.Code = int32(e.Code)
	grpcStatus.Message = e.Message
	grpcStatus.Details = make([]*any.Any, 0)
	anyObj := new(any.Any)
	if e.Service == "" {
		anyObj.TypeUrl = ServiceName
	} else {
		anyObj.TypeUrl = e.Service
	}
	grpcStatus.Details = append(grpcStatus.Details, anyObj)
	e.s = grpcStatus
}

// New create a new pointer of Error
func New(code ErrorCode, messages ...string) error {
	e := new(Error)
	e.Code = code
	e.Service = os.Getenv("service")
	if e.Service == "" {
		e.Service = ServiceName
	}

	if len(messages) == 0 {
		e.Message = errorMessages[code]
	} else {
		for idx, msg := range messages {
			if idx == 0 {
				e.Message = msg
			} else {
				e.Message = msg + errMsgSep + e.Message
			}
		}
	}
	e.grpcEncode()

	return status.FromProto(e.s).Err()
}

// FromError : convert error to *Error
func FromError(err error) *Error {
	if _, ok := err.(*Error); ok {
		return err.(*Error)
	}

	e := new(Error)

	st, ok := status.FromError(err)
	if !ok {
		// Error was not a status error
		e.Code = ErrorCode(ECUnknown)
		e.Message = err.Error()
		e.Service = ServiceName
	} else {
		e.Code = ErrorCode(st.Code())
		e.Message = st.Message()
		if st.Proto() != nil && len(st.Proto().GetDetails()) > 0 {
			anyObj := st.Proto().GetDetails()[0]
			e.Service = anyObj.GetTypeUrl()
		} else {
			e.Service = ServiceName
		}
	}

	return e
}

// WithMessage : annotates error with a new message
func WithMessage(err error, messages ...string) error {
	e := FromError(err)
	for _, msg := range messages {
		e.Message = msg + errMsgSep + e.Message
	}
	e.grpcEncode()
	return status.FromProto(e.s).Err()
}
