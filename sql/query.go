package sql

func Like(value string) string {
	return "%" + value + "%"
}

func LLike(value string) string {
	return "%" + value
}

func RLike(value string) string {
	return value + "%"
}
