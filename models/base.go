package models

/**
 * @author: zhashut
 * Date: 2025/4/19
 * Time: 0:27
 * Description: 基础结构体
 */

type PageRequest struct {
	Page      int64  `json:"page"`      // 当前页号
	PageSize  int64  `json:"pageSize"`  // 页面大小
	SortField string `json:"sortField"` // 排序字段
	SortOrder string `json:"sortOrder"` // 排序顺序(默认升序)
}

// OnlyIDRequest 传递单个 id 的请求-通用
type OnlyIDRequest struct {
	ID int64 `form:"id" json:"id"`
}
