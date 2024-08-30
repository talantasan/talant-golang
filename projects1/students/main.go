package main

import "fmt"

// import "fmt"

func main() {
	students := newStudents()

	// students.print()

	cls1, cls2 := splitClass(students, 4)
	cls1.print()
	fmt.Println("'-----------------'")
	cls2.print()

}