package errors

// BadRequest new BadRequest error
func BadRequest(reason string) *Error {
	return New(400, reason)
}

// IsBadRequest determines if err is BadRequest error.
func IsBadRequest(err *Error) bool {
	return err.Code == 400
}

// Unauthorized new Unauthorized error
func Unauthorized(reason string) *Error {
	return New(401, reason)
}

// IsUnauthorized determines if err is Unauthorized error.
func IsUnauthorized(err *Error) bool {
	return err.Code == 401
}

// Forbidden new Forbidden error
func Forbidden(reason string) *Error {
	return New(403, reason)
}

// IsForbidden determines if err is Forbidden error.
func IsForbidden(err *Error) bool {
	return err.Code == 403
}

// NotFound new NotFound error
func NotFound(reason string) *Error {
	return New(404, reason)
}

// IsNotFound determines if err is NotFound error.
func IsNotFound(err *Error) bool {
	return err.Code == 404
}

// Conflict new Conflict error
func Conflict(reason string) *Error {
	return New(409, reason)
}

// IsConflict determines if err is Conflict error.
func IsConflict(err *Error) bool {
	return err.Code == 409
}

// InternalServer new InternalServer error
func InternalServer(reason string) *Error {
	return New(500, reason)
}

// IsInternalServer determines if err is InternalServer error.
func IsInternalServer(err *Error) bool {
	return err.Code == 500
}

// ServiceUnavailable new ServiceUnavailable error
func ServiceUnavailable(reason string) *Error {
	return New(503, reason)
}

// IsServiceUnavailable determines if err is ServiceUnavailable error.
func IsServiceUnavailable(err *Error) bool {
	return err.Code == 503
}

// GatewayTimeout new GatewayTimeout error
func GatewayTimeout(reason string) *Error {
	return New(504, reason)
}

// IsGatewayTimeout determines if err is GatewayTimeout error.
func IsGatewayTimeout(err *Error) bool {
	return err.Code == 504
}

// ClientClosed new ClientClosed error
func ClientClosed(reason string) *Error {
	return New(499, reason)
}

// IsClientClosed determines if err is ClientClosed error.
func IsClientClosed(err *Error) bool {
	return err.Code == 499
}
