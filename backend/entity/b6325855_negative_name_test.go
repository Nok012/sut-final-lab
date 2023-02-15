package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"

	. "github.com/onsi/gomega"
)

func TestNameNotBlank(t *testing.T) {

	g := NewGomegaWithT(t)

	c := Customer{
		Name:       "",
		Email:      "Nok12@gamil.com",
		CustomerID: "L1234567",
	}

	ok, err := govalidator.ValidateStruct(c)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).NotTo(BeNil())

	g.Expect(err.Error()).To(Equal("Name: non zero value required"))

}
