package e

var MsgFlags = map[int]string{
	// 200
	OK:                              "200.OK",


	// 400系エラーメッセージ
	BAD_REQUEST:                     "400.リクエストが不正です",
	JSON_PARSE_ERROR:                "400.リクエストの形式が不正です",
	INVALID_ID:                      "400.無効なIDを使用しています",
	CASE_NOT_OPEN:                   "400.この案件は募集を終了しています",
	ALREADY_APPLIED:                 "400.既にこの案件に応募しています",

	// 403系エラーメッセージ
	FORBIDDEN:                       "403.この操作を行う権限がありません",
	FORBIDDEN_UPDATE_CASE:           "403.他のユーザーの案件を編集する権限がありません",


	// 404系エラーメッセージ
	NOT_FOUND:                       "404.Not Found",
	NOT_FOUND_USER:                  "404.ユーザーが見つかりません",
	NOT_FOUND_CASE:                  "404.案件が見つかりません",
	NOT_FOUND_APPLICATION:           "404.応募が見つかりません",
	NOT_FOUND_REVIEW:                "404.レビューが見つかりません",


	// 500系エラーメッセージ
	SERVER_ERROR:                    "500.サーバーエラー",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[SERVER_ERROR]
}