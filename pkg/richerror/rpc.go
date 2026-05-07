package richerror

import (
	"fmt"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *RichError) ToRPCError() error {
	grpcCode := mapKindToGRPCCode(r.kind)

	st := status.New(grpcCode, r.message)

	if len(r.fields) > 0 {
		br := &errdetails.BadRequest{}

		for field, msg := range r.fields {
			br.FieldViolations = append(br.FieldViolations, &errdetails.BadRequest_FieldViolation{
				Field:       field,
				Description: msg,
			})
		}

		detailed, err := st.WithDetails(br)
		if err == nil {
			st = detailed
		}
	}

	if len(r.meta) > 0 {
		md := &errdetails.ErrorInfo{
			Reason:   "RICH_ERROR",
			Domain:   string(r.operation),
			Metadata: convertMetaToStringMap(r.meta),
		}

		detailed, err := st.WithDetails(md)
		if err == nil {
			st = detailed
		}
	}

	return st.Err()
}

func convertMetaToStringMap(m map[string]interface{}) map[string]string {
	out := map[string]string{}
	for k, v := range m {
		out[k] = fmt.Sprint(v)
	}
	return out
}

func mapKindToGRPCCode(kind Kind) codes.Code {
	switch kind {
	case KindInvalid:
		return codes.InvalidArgument
	case KindNotFound:
		return codes.NotFound
	case KindUnexpected:
		return codes.Internal
	case KindForbidden:
		return codes.PermissionDenied
	default:
		return codes.Unknown
	}
}
