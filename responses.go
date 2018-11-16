package main

//定义返回json的格式
type JsonResponse struct {
	// Reserved field to add some meta information to the API response
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

type JsonErrorResponse struct {
	Error *ApiError `json:"error"`
}
//定义失败
type ApiError struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
}
