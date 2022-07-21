package types

// BaseResponse base
type BaseResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

// BaseQueryRequest base query
type BaseQueryRequest struct {
	Page int `form:"page"` // 必传 页码，min=1
	Size int `form:"size"` // 必传，单页size，min=1，max=100
}
