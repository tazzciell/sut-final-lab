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
		Salary:       16000,
		EmployeeCode: "HR-1024",
	}

	ok, err := employee.Validate()

	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())

}
