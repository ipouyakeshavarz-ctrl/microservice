package grpcerror

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Extract(err error) (string, map[string]string, codes.Code) {
	st, ok := status.FromError(err)
	if !ok {
		return err.Error(), nil, codes.Internal
	}

	fields := map[string]string{}

	for _, detail := range st.Details() {

		switch d := detail.(type) {

		case *errdetails.BadRequest:
			for _, v := range d.FieldViolations {
				fields[v.Field] = v.Description
			}
		}
	}

	return st.Message(), fields, st.Code()
}
