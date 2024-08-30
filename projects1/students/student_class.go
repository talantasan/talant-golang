package main

import "fmt"

type students []string

func newStudents() students {
	newStd := students{}

	student_list := students{"Student-1", "Student-2", "Student-3"}
	class_list := students{"Math", "Biology", "History"}

	for _, std := range student_list{
		for _, clss := range class_list{
			newStd = append(newStd, std+" of "+clss)
		}
	}

	return newStd
}

func (s students) print(){
	for i, stds := range s{
		fmt.Println(i, stds)
	}
}

func splitClass(st students, gr int) (students, students){
	return st[:gr], st[gr:]
}

