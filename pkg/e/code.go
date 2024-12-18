package e

// これをHTTPステータスコードとして使うわけではない
const (
	// 200系
	OK         = 200
	CREATED    = 201
	NO_CONTENT = 204

	// 400系
	BAD_REQUEST                = 400
	JSON_PARSE_ERROR           = 400001
	INVALID_UUID               = 400002
	CASE_NOT_OPEN              = 400003
	OVER_CONFIRMATION_DEADLINE = 400004
	INVALID_USER_ID            = 400005
	MATCHING_NOT_COMPLETED     = 400006
	INVALID_PASSWORD           = 400007

	// 401系
	UNAUTHORIZED              = 401
	AUTH_REQUIRED             = 401001
	INVALID_TOKEN             = 401002
	INVALID_TOKEN_FORMAT      = 401003
	INVALID_EMAIL_OR_PASSWORD = 401004
	SESSION_EXPIRED           = 401005
	INVALID_DEVICE            = 401006
	INVALID_SESSION_ID        = 401007
	TOKEN_EXPIRED             = 401008

	// 403系
	FORBIDDEN                    = 403
	FORBIDDEN_UPDATE_APPLICATION = 403001
	FORBIDDEN_DELETE_CASE        = 403002
	FORBIDDEN_REVIEW             = 403003

	// 404系
	NOT_FOUND             = 404
	NOT_FOUND_USER        = 404001
	NOT_FOUND_CASE        = 404002
	NOT_FOUND_APPLICATION = 404003
	NOT_FOUND_REVIEW      = 404004
	NOT_FOUND_MATCHING    = 404005
	NOT_FOUND_SESSION     = 404006

	// 409系
	ALREADY_APPLIED          = 409001
	ALREADY_REGISTERED_EMAIL = 409002
	ALREADY_REVIEWED         = 409003

	// 500系
	SERVER_ERROR = 500
)
