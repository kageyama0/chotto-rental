package e

var MsgFlags = map[int]string{
	// 200
	OK:                              "OK",
	CREATED:                         "Created",
	NO_CONTENT:                      "No Content",

	// 400系エラーメッセージ
	BAD_REQUEST:                     "リクエストが不正です",
	JSON_PARSE_ERROR:                "リクエストの形式が不正です",
	INVALID_PARAMS:                      "無効なIDを使用しています",
	CASE_NOT_OPEN:                   "この案件は募集を終了しています",
	OVER_CONFIRMATION_DEADLINE:      "確認期限が過ぎています",

	// 401系エラーメッセージ
	UNAUTHORIZED:                    "認証エラー",
	INVALID_TOKEN:                   "無効なトークンです",
	INVALID_TOKEN_FORMAT:            "無効なトークンフォーマットです",
	AUTH_REQUIRED:                   "認証情報が必要です",
	INVALID_EMAIL_OR_PASSWORD:       "メールアドレスまたはパスワードが間違っています",

	// 403系エラーメッセージ
	FORBIDDEN:                       "この操作を行う権限がありません",
	FORBIDDEN_UPDATE_APPLICATION:    "この応募のステータスを更新する権限がありません",
	FORBIDDEN_DELETE_CASE:           "この案件を削除する権限がありません",


	// 404系エラーメッセージ
	NOT_FOUND:                       "Not Found",
	NOT_FOUND_USER:                  "ユーザーが見つかりません",
	NOT_FOUND_CASE:                  "案件が見つかりません",
	NOT_FOUND_APPLICATION:           "応募が見つかりません",
	NOT_FOUND_REVIEW:                "レビューが見つかりません",

	// 409系エラーメッセージ
	ALREADY_APPLIED:                 "既にこの案件に応募しています",
	EMAIL_ALREADY_EXISTS:            "既にこのメールアドレスは登録されています",


	// 500系エラーメッセージ
	SERVER_ERROR:                    "サーバーエラー",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[SERVER_ERROR]
}
