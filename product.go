package entities

type Product struct {
	ID         uint   "json: 'id'"
	Firstname  string "json: 'fname'"
	Lastname   string "json: 'lname'"
	Department string "json: 'department'"
	JoinDate   string "json: 'jdate'"
}
