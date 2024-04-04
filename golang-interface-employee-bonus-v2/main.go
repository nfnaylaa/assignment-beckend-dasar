package main

import "fmt"

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name         string
	BaseSalary   int
	WorkingMonth int
}

type Senior struct {
	Name            string
	BaseSalary      int
	WorkingMonth    int
	PerformanceRate float64
}

type Manager struct {
	Name             string
	BaseSalary       int
	WorkingMonth     int
	PerformanceRate  float64
	BonusManagerRate float64
}

func (j Junior) GetBonus() float64 {
	prorata := float64(j.WorkingMonth) / 12
	if prorata > 1 {
		prorata = 1
	}
	return float64(j.BaseSalary) * prorata
}

func (s Senior) GetBonus() float64 {
	prorata := float64(s.WorkingMonth) / 12
	if prorata > 1 {
		prorata = 1
	}
	return (2 * float64(s.BaseSalary) * prorata) + (float64(s.BaseSalary) * s.PerformanceRate)
}

func (m Manager) GetBonus() float64 {
	prorata := float64(m.WorkingMonth) / 12
	if prorata > 1 {
		prorata = 1
	}
	return (2 * float64(m.BaseSalary) * prorata) + (m.PerformanceRate * float64(m.BaseSalary)) + (m.BonusManagerRate * float64(m.BaseSalary))
}
func EmployeeBonus(employee Employee) float64 {

	return employee.GetBonus() // TODO: replace this
}

func TotalEmployeeBonus(employees []Employee) float64 {
	totalBonus := 0.0
	for _, employee := range employees {
		totalBonus += employee.GetBonus()
	}
	return totalBonus // TODO: replace this
}

func main() {
	employees := []Employee{
		Junior{Name: "Junior A", BaseSalary: 100000, WorkingMonth: 12},
		Junior{Name: "Junior B", BaseSalary: 100000, WorkingMonth: 12},
		Junior{Name: "Junior C", BaseSalary: 100000, WorkingMonth: 12},
	}
	fmt.Println(TotalEmployeeBonus(employees))
}
