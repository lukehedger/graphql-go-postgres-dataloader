package api

type Resolver struct {
	DB *DB
}

type Person struct {
	ID   string
	Name string
}

type PersonResolver struct {
	Person *Person
}
