package persistence

func GetAccessToken() (string, error) {
	return readFile("access.token")
}

func SetAccessToken(token string) error {
	return writeFile("access.token", token)
}

func GetRefreshToken() (string, error) {
	return readFile("refresh.token")
}

func SetRefreshToken(token string) error {
	return writeFile("refresh.token", token)
}
