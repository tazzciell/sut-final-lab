package test

import (
	. "github.com/onsi/gomega"
	"test/entity"
	"testing"
)

func TestTask2(t *testing.T) {
	g := NewGomegaWithT(t)

	employee := entity.Employees{
		Name:         "Tan",
		Salary:       12000,
		EmployeeCode: "HR-1024",
	}

	ok, err := employee.Validate()

	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(ContainSubstring("Salary must be between 15000 and 200000"))
}