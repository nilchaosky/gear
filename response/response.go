package response

type (
	response[T any] struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    T      `json:"data"`
	}

	Page[T any] struct {
		Current int   `json:"current"`
		Size    int   `json:"size"`
		Total   int64 `json:"total"`
		Records []T   `json:"records"`
	}
)

func Success[T any](data T) *response[T] {
	return &response[T]{
		Code:    0,
		Message: "success",
		Data:    data,
	}
}

func SuccessNull() *response[any] {
	return &response[any]{
		Code:    0,
		Message: "success",
	}
}

func SuccessPage[T any](page *Page[T]) *response[*Page[T]] {
	return &response[*Page[T]]{
		Code:    0,
		Message: "success",
		Data:    page,
	}
}

func Fail(message string) *response[any] {
	return &response[any]{
		Code:    -1,
		Message: message,
		Data:    nil,
	}
}
