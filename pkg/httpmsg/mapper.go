package httpmsg

import "myapp/pkg/grpcerror"

func Error(err error) (map[string]interface{}, int) {

	msg, fields := grpcerror.Extract(err)

	if len(fields) > 0 {

		return map[string]interface{}{
			"message": msg,
			"errors":  fields,
		}, 400
	}

	return map[string]interface{}{
		"message": msg,
	}, 400
}

//func mapKindToHTTPStatusCode(kind richerror.Kind) int {
//	switch kind {
//	case richerror.KindInvalid:
//		return http.StatusUnprocessableEntity // 422
//	case richerror.KindNotFound:
//		return http.StatusNotFound // 404
//	case richerror.KindForbidden:
//		return http.StatusForbidden // 403
//	case richerror.KindUnexpected:
//		return http.StatusInternalServerError
//	default:
//		return http.StatusBadRequest
//	}
//}
