package utils

type AlreadyFind struct{}

func (AlreadyFind) Error() string {
	return "already present in database"
}
