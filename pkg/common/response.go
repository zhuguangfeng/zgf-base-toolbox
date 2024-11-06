package common

type Response struct {
	StatusCode int         `json:"statusCode" example:"900"`                           // 操作状态码：成功200，失败-1
	Error      string      `json:"error,omitempty" example:"GoChat.User.UnAuthorized"` // 业务状态码，如果操作状态码为200，则为成功业务状态码，如果操作状态码为-1，则为失败业务状态码
	Message    string      `json:"message,omitempty" example:"没有任务操作权限"`       // 业务状态消息，如果操作状态码为200，则为成功业务消息，如果操作状态码为-1，则为失败业务消息
	ReturnObj  interface{} `json:"returnObj,omitempty"`                                // 业务数据
}

type ListObj struct {
	CurrentCount int         `json:"currentCount,omitempty" example:"10"` // 本次数据条数
	TotalCount   int         `json:"totalCount,omitempty" example:"28"`   // 总数据条数
	TotalPage    int         `json:"totalPage,omitempty" example:"3"`     // 总数据页数
	Result       interface{} `json:"result"`                              // 业务数据
}
