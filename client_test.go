package midtrans_test

import (
    "testing"
    "github.com/cheekybits/is"
    "midtrans"
)

func TestDefaultEnvironmentType(t *testing.T) {
    is := is.New(t)

    midclient := midtrans.NewClient()
    is.Equal(midtrans.Sandbox, midclient.ApiEnvType)
}