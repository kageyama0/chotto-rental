package e

// これをHTTPステータスコードとして使うわけではない
const (
	// 200系
	OK                       = 200
	CREATED 				         = 201
	NO_CONTENT               = 204

	// 400系
	BAD_REQUEST              = 400
	JSON_PARSE_ERROR         = 400001
	INVALID_ID			         = 400002
	CASE_NOT_OPEN						 = 400003
	ALREADY_APPLIED					 = 400004

	// 401系
	UNAUTHORIZED             = 401


	// 403系
	FORBIDDEN                = 403
	FORBIDDEN_UPDATE_CASE    = 403001

	// 404系
	NOT_FOUND				         = 404
	NOT_FOUND_USER           = 404001
	NOT_FOUND_CASE           = 404002
	NOT_FOUND_APPLICATION    = 404003
	NOT_FOUND_REVIEW         = 404004


	// 500系
	SERVER_ERROR             = 500
)
