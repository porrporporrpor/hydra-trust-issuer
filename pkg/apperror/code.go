package apperror

type ErrorCode int
type ErrorMessage string

const (
	// BadRequestError ... Missing mandatory parameter or Invalid data type
	BadRequestError ErrorCode = 4001

	// UnAuthorizedError ... Failed from authorization provider
	UnAuthorizedError ErrorCode = 4011

	// ForbiddenError ... Failed from access to forbidden resources
	ForbiddenError ErrorCode = 4031

	// NotFoundError ... Failed from access to unknown resources
	NotFoundError     ErrorCode = 4041
	PageNotFoundError ErrorCode = 4042

	// ConflictError ... Failed when interacting with database
	ConflictError ErrorCode = 4091

	// GoneError ... error code to use when request was used
	GoneError ErrorCode = 4101

	// InternalServerError ... Failed from access to server
	InternalServerError ErrorCode = 5001
	ExternalCallError   ErrorCode = 5002
)

var (
	Message = map[ErrorCode]ErrorMessage{
		// 400x
		BadRequestError: "Bad Request",

		// 401x
		UnAuthorizedError: "Unauthorized",

		// 403x
		ForbiddenError: "Forbidden",

		// 404x
		NotFoundError:     "Not Found",
		PageNotFoundError: "Page Not Found",

		// 409x
		ConflictError: "Conflict",

		// 410x
		GoneError: "Gone",

		// 500x
		InternalServerError: "Internal Server Error",
		ExternalCallError:   "External Server Error",
	}
)
