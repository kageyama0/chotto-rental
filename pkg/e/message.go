package e

var MsgFlags = map[int]string{
	// 200
	OK:         "OK",
	CREATED:    "Created",
	NO_CONTENT: "No Content",

	// 400系エラーメッセージ
	BAD_REQUEST:                "リクエストが不正です",
	JSON_PARSE_ERROR:           "リクエストの形式が不正です",
	INVALID_UUID:               "無効なIDを使用しています",
	CASE_NOT_OPEN:              "この案件は募集を終了しています",
	OVER_CONFIRMATION_DEADLINE: "確認期限が過ぎています",
	INVALID_USER_ID:            "無効なユーザーIDを使用しています",
	MATCHING_NOT_COMPLETED:     "完了していないマッチングにはレビューできません",
	INVALID_PASSWORD:           "パスワードが間違っています",

	// 401系エラーメッセージ
	UNAUTHORIZED:              "認証エラー",
	INVALID_TOKEN:             "無効なトークンです",
	INVALID_TOKEN_FORMAT:      "無効なトークンフォーマットです",
	AUTH_REQUIRED:             "認証情報が必要です",
	INVALID_EMAIL_OR_PASSWORD: "メールアドレスまたはパスワードが間違っています",
	SESSION_EXPIRED:           "セッションが期限切れです",
	INVALID_DEVICE:            "無効なデバイスです",
	INVALID_SESSION_ID:        "無効なセッションIDを使用しています",
	INVALID_SESSION:           "無効なセッションです",

	// 403系エラーメッセージ
	FORBIDDEN:                    "この操作を行う権限がありません",
	FORBIDDEN_UPDATE_APPLICATION: "この応募のステータスを更新する権限がありません",
	FORBIDDEN_DELETE_CASE:        "この案件を削除する権限がありません",
	FORBIDDEN_REVIEW:             "このマッチングにレビューを投稿する権限がありません",

	// 404系エラーメッセージ
	NOT_FOUND:             "Not Found",
	NOT_FOUND_USER:        "ユーザーが見つかりません",
	NOT_FOUND_CASE:        "案件が見つかりません",
	NOT_FOUND_APPLICATION: "応募が見つかりません",
	NOT_FOUND_REVIEW:      "レビューが見つかりません",
	NOT_FOUND_MATCHING:    "マッチングが見つかりません",
	NOT_FOUND_SESSION:     "セッションが見つかりません",
	NOT_FOUND_SESSION_ID:  "セッションIDが見つかりません",

	// 409系エラーメッセージ
	ALREADY_APPLIED:          "既にこの案件に応募しています",
	ALREADY_REGISTERED_EMAIL: "既にこのメールアドレスは登録されています",
	ALREADY_REVIEWED:         "既にこのマッチングにレビューを投稿済みです",

	// 500系エラーメッセージ
	SERVER_ERROR: "サーバーエラー",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[SERVER_ERROR]
}
