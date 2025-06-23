package student

import "fmt"

type Student struct {
	ID     int
	Name   string
	Age    int
	Email  string
	Grades []int
}

func (s Student) GetAverageGrades() float64 {
	len := len(s.Grades)
	var sum int
	for _, val := range s.Grades {
		sum = sum + val
	}
	avg := float64(sum) / float64(len)
	return avg
}

func (s Student) DisplayInfo() {
	fmt.Println("📚 Student Information:")

	fmt.Println("--------------------------")

	fmt.Println("🆔 ID     :", s.ID)
	fmt.Println("👤 Name   :", s.Name)
	fmt.Println("🎂 Age    :", s.Age)
	fmt.Println("📧 Email  :", s.Email)
	fmt.Println("📊 Grades :", s.Grades)
	fmt.Printf("📈 Average Grade: %.2f\n", s.GetAverageGrades())
}
