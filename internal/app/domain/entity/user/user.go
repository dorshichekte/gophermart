package entity

func NewUser(login, password string) User {
	return User{
		Login:        login,
		PasswordHash: password,
	}
}
