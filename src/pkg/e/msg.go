package e

var MsgFlags = map[int]string{
	SUCCESS:               "ok",
	ERROR:                 "fail",
	INVALID_PARAMS:        "请求参数错误",
	ERROR_EXIST_TAG_FAIL:  "获取已存在标签失败",
	ERROR_NOT_EXIST_TAG:   "该标签不存在",
	ERROR_GET_NAMES_FAIL:  "获取名字列表失败",
	ERROR_COUNT_TAG_FAIL:  "统计标签失败",
	ERROR_ADD_TAG_FAIL:    "新增标签失败",
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
