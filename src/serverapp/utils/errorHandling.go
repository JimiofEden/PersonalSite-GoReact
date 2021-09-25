package utils

func CheckAndHandleErr(err error) {
	if err != nil {
		panic(err)
	}
}