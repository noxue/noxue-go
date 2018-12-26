package utils

type Error string

func (e Error) Error() string {
	return string(e)
}

func CheckErr(err error) {
	if err != nil {
		panic(Error(err.Error()))
	}
}
