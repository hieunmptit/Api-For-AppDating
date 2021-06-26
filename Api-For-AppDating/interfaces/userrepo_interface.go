package interfaces

type ResponseUser struct {
	ID       uint
	Username string
	Email    string
	Profile  []ResponseProfile
}
