package utils

type AlreadyFind struct{}

func (AlreadyFind) Error() string {
	return "already present in database"
}

type NotFound struct{}

func (NotFound) Error() string {
	return "Not present in database"
}