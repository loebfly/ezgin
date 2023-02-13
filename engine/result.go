package engine

type Result[D any] struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
	Data    D      `json:"data"`
	Page    *Page  `json:"page"`
}

// ToAnyRes 转换为Result[any]
func (receiver Result[D]) ToAnyRes() Result[any] {
	if receiver.Status != 1 {
		return Result[any]{
			Status:  receiver.Status,
			Message: receiver.Message,
		}
	}
	return Result[any]{
		Status:  receiver.Status,
		Message: receiver.Message,
		Data:    receiver.Data,
		Page:    receiver.Page,
	}
}

// ConvResDataType 自由将Result[From]转换为Result[To]
func ConvResDataType[From, To any](from Result[From]) Result[To] {
	res := from.ToAnyRes()
	if res.Data == nil {
		return Result[To]{
			Status:  res.Status,
			Message: res.Message,
		}
	}
	return Result[To]{
		Status:  res.Status,
		Message: res.Message,
		Data:    res.Data.(To),
		Page:    res.Page,
	}
}

type Page struct {
	Count int `json:"count"`
	Index int `json:"index"`
	Size  int `json:"size"`
	Total int `json:"total"`
}

func ErrorRes(status int, message string) Result[any] {
	return Result[any]{
		Status:  status,
		Message: message,
	}
}

func SuccessRes[D any](data D, message ...string) Result[D] {
	msg := "success"
	if len(message) > 0 {
		msg = message[0]
	}
	return Result[D]{
		Status:  1,
		Message: msg,
		Data:    data,
	}
}

func SuccessPageRes[D any](data D, page Page, message ...string) Result[D] {
	msg := "success"
	if len(message) > 0 {
		msg = message[0]
	}
	return Result[D]{
		Status:  1,
		Message: msg,
		Data:    data,
		Page:    &page,
	}
}
