package grpcerror

import (
	"myapp/pkg/errmsg"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MapGRPCError(err error) (string, int, bool) {
	st, ok := status.FromError(err)
	if !ok {
		return "", 0, false
	}

	msg := st.Message()

	switch st.Code() {

	case codes.InvalidArgument:
		return msg, http.StatusBadRequest, true

	case codes.NotFound:
		return msg, http.StatusNotFound, true

	case codes.Unauthenticated:
		return msg, http.StatusUnauthorized, true

	case codes.PermissionDenied:
		return msg, http.StatusForbidden, true

	case codes.AlreadyExists:
		return msg, http.StatusConflict, true

	case codes.Internal:
		return errmsg.ErrorMsgSomethingWentWrong, http.StatusInternalServerError, true
	}

	return msg, http.StatusInternalServerError, true
}
