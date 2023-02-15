package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"

	. "github.com/onsi/gomega"
)

func TestCustomerValid(t *testing.T) {

	g := NewGomegaWithT(t)

	c := Customer{
		Name:       "Name",
		Email:      "Nok12@gamil.com",
		CustomerID: "L1234567",
	}

	ok, err := govalidator.ValidateStruct(c)

	g.Expect(ok).To(BeTrue())

	g.Expect(err).To(BeNil())

}
