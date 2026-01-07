package test

import (
	. "github.com/onsi/gomega"
	"test/entity"
	"testing"
)

func TestTask3(t *testing.T) {
	g := NewGomegaWithT(t)

	employee := entity.Employees{
		Name:         "Tan",
		Salary:       12000,
		EmployeeCode: "HR-10245656",
	}

	ok, err := employee.Validate()

	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(ContainSubstring("EmployeeCode must be 2 uppercase English letters (A-Z) followed by ‘-’ and 4 digits (0-9)"))
}