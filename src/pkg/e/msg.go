package e

var MsgFlags = map[int]string{
	SUCCESS:               "ok",
	ERROR:                 "fail",
	InvalidParams:         "请求参数错误",
	ErrorExistTagFail:     "获取已存在标签失败",
	ErrorNotExistTag:      "该标签不存在",
	ErrorGetNamesFail:     "获取名字列表失败",
	ErrorCountTagFail:     "统计标签失败",
	ErrorGetWxOpenFail:    "获取微信openid失败",
	ERROR_EDIT_TAG_FAIL:   "修改标签失败",
	ERROR_DELETE_TAG_FAIL: "删除标签失败",
	ERROR_EXPORT_TAG_FAIL: "导出标签失败",
	ERROR_IMPORT_TAG_FAIL: "导入标签失败",
	ERROR_QUERY_DB_FAIL:   "查询数据库失败",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
