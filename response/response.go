package response

type (
	Response[T any] struct {
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

func Success[T any](data T) *Response[T] {
	return &Response[T]{
		Code:    0,
		Message: "success",
		Data:    data,
	}
}

func SuccessNull() *Response[any] {
	return &Response[any]{
		Code:    0,
		Message: "success",
	}
}

func SuccessPage[T any](page *Page[T]) *Response[*Page[T]] {
	return &Response[*Page[T]]{
		Code:    0,
		Message: "success",
		Data:    page,
	}
}

func Fail(message string) *Response[any] {
	return &Response[any]{
		Code:    -1,
		Message: message,
		Data:    nil,
	}
}
