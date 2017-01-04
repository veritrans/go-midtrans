package midtrans

import (
    "testing"
    "github.com/cheekybits/is"
)

func TestDefaultEnvironmentType(t *testing.T) {
    is := is.New(t)
    is.Equal(Sandbox, ClientSetting.ApiEnvType)
}