package domain

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Age       int
	IsMarried bool
	Password  string
}

func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}
