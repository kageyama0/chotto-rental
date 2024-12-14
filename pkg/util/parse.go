package util

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/pkg/e"
)

func parseUUID(input string) (uid *uuid.UUID, MsgCode int) {
    parsedUUID, err := uuid.Parse(input)
    if err != nil {
        return nil, e.INVALID_UUID
    }
    return &parsedUUID, e.OK
}

func getUserID(c *gin.Context) (userID *uuid.UUID, MsgCode int) {
    userIDRaw, exists := c.Get("userID")
    if !exists {
        return nil, e.INVALID_USER_ID
    }

    userIDStr, ok := userIDRaw.(string)
    if !ok {
        return nil, e.INVALID_USER_ID
    }

    return parseUUID(userIDStr)
}


// -- GetParams: APIの最初の処理として、URLのパラメータとユーザーIDを取得します。
func GetParams(c *gin.Context, params []string) (map[string]uuid.UUID, *uuid.UUID, int) {
	// URLのパラメータの取得
	parsedParams := make(map[string]uuid.UUID)
	for _, param := range params {
		parsedParam, MsgCode := parseUUID(c.Param(param))
		if MsgCode != e.OK {
			return nil, nil, MsgCode
		}
		parsedParams[param] = *parsedParam
	}

	// ユーザーIDの取得
	userID, MsgCode := getUserID(c)
	if MsgCode != e.OK {
		return nil, nil, MsgCode
	}

	return parsedParams, userID, e.OK
}

// -- parseJson
