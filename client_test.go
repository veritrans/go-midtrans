package midtrans_test

import (
	"testing"

	midtrans "github.com/veritrans/go-midtrans"

	"github.com/cheekybits/is"
)

func TestDefaultEnvironmentType(t *testing.T) {
	is := is.New(t)

	midclient := midtrans.NewClient()
	is.Equal(midtrans.Sandbox, midclient.APIEnvType)
}
