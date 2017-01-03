package midtrans_test

import (
    "github.com/cheekybits/is"
    "testing"
    "midtrans"
)

func TestEnvironmentType(t *testing.T) {
    is := is.New(t)
    is.Equal("https://api.sandbox.veritrans.co.id/v2", midtrans.Sandbox.String())
    is.Equal("https://api.veritrans.co.id/v2", midtrans.Production.String())
}
