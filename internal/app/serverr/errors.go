package serverr

type UserNotFound struct {
}

func (*UserNotFound) Error() string {
	return "User not found"
}
