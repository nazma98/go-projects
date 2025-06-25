package student

var Students []Student

func AddStudent(s Student) {
	Students = append(Students, s)
}

func DeleteStudent(id int) {
	for ind, val := range Students {
		if id == val.ID {
			Students = append(Students[:ind], Students[ind+1:]...)
			return
		}
	}
}

func PrintAllStudents() {
	for _, val := range Students {
		val.DisplayInfo()
	}
}
