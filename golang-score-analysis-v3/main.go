package main

import (
	"fmt"
	"sort"
)

type School struct {
	Name    string
	Address string
	Grades  []int
}

func (s *School) AddGrade(grades ...int) {
	s.Grades = append(s.Grades, grades...)
}

func Analysis(s School) (float64, int, int) {

	if len(s.Grades) == 0 {
		return 0, 0, 0
	}
	total := 0
	for _, grade := range s.Grades {
		total += grade
	}
	average := float64(total) / float64(len(s.Grades))
	sort.Ints(s.Grades)
	minimum := s.Grades[0]
	maximal := s.Grades[len(s.Grades)-1]

	return average, minimum, maximal // TODO: replace this
}

// gunakan untuk melakukan debugging
func main() {
	avg, min, max := Analysis(School{
		Name:    "Imam Assidiqi School",
		Address: "Jl. Imam Assidiqi",
		Grades:  []int{100, 90, 80, 70, 60},
	})

	fmt.Println(avg, min, max)
}
