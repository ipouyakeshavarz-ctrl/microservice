package httpmsg

import (
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/grpcerror"
	"net/http"

	"google.golang.org/grpc/codes"
)

func Error(err error) (dto.ErrorResponse, int) {

	msg, fields, code := grpcerror.Extract(err)

	status := mapCodeToHTTPStatus(code)

	if len(fields) > 0 {
		return dto.ErrorResponse{
			Message: msg,
			Errors:  fields,
		}, status
	}

	return dto.ErrorResponse{
		Message: msg,
	}, status
}

func mapCodeToHTTPStatus(code codes.Code) int {

	switch code {

	case codes.InvalidArgument:
		return http.StatusUnprocessableEntity

	case codes.NotFound:
		return http.StatusNotFound

	case codes.PermissionDenied:
		return http.StatusForbidden

	case codes.Unauthenticated:
		return http.StatusUnauthorized

	case codes.AlreadyExists:
		return http.StatusConflict

	case codes.Internal:
		return http.StatusInternalServerError

	default:
		return http.StatusBadRequest
	}
}
