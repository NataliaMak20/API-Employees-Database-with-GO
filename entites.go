package data

type Employee struct {
	Id   int
	Age  int
	City string
	Name string
}

func IsCool(emp Employee) bool {
	if emp.Name == "Natalia" {
		return true
	}
	return false
}
