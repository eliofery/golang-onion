package datastruct

const UserTableName = "users"

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
	Password  string
}
