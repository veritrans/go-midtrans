package midtrans_test

import (
	"testing"

	"github.com/cheekybits/is"
	midtrans "github.com/veritrans/go-midtrans"
)

func TestEnvironmentType(t *testing.T) {
	is := is.New(t)
	is.Equal("https://api.sandbox.midtrans.com", midtrans.Sandbox.String())
	is.Equal("https://api.midtrans.com", midtrans.Production.String())
}
