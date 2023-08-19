
syntax = "{{.Syntax}}"

info (
	title: "view"
	desc: "抽取类型,方便在其他api使用"
	author: "{{.Author}}"
	email: "{{.Email}}"
	version: "{{.Version}}"
)

// 基本查询参数, 根据自己需要进行修改
type SearchBase {
    Keyword string `json:"keyword,optional"`          // 关键字
    Cursor int64 `json:"cursor,optional"`             // 游标,基于游标分页时使用
    CursorAsc bool `json:"cursorAsc,optional"`          // 游标分页时方向 true:asc  false:desc
    PageSize int `json:"pageSize,default=20,optional"`           // 每页条数
    PageCurrent int `json:"pageCurrent,default=1,optional"`     // 当前页
    OrderSort []string `json:"orderSort,optional"`    // 排序 eg： ["create_time asc", "id desc"]
    SearchPlus string `json:"searchPlus,optional"` // 加强版自定义搜索参数
}

// 统一分页返回
type PageBase {
    Total int64 `json:"total,omitempty"` // 总条数
    PageCurrent int `json:"pageCurrent,omitempty"` // 当前页
    PageSize int `json:"pageSize,omitempty"` // 分页大小
    PageTotal int `json:"pageTotal,omitempty"` // 总分页数
    LastCursor int64 `json:"lastCursor,omitempty"` // 使用游标分页时, 返回最后一个游标
}