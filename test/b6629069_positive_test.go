package test

import (
	. "github.com/onsi/gomega"
	"test/entity"
	"testing"
)

func TestTask1(t *testing.T) {
	g := NewGomegaWithT(t)

	employee := entity.Employees{
		Name:         "tanthorsang",
		Salary:       20000,
		EmployeeCode: "AB-1234",
	}

	ok, err := employee.Validate()

	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())

}
