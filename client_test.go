package midtrans_test

import (
	"testing"

	"time"

	"strconv"

	midtrans "github.com/veritrans/go-midtrans"

	"github.com/cheekybits/is"
)

func TestDefaultEnvironmentType(t *testing.T) {
	is := is.New(t)

	midclient := midtrans.NewClient()
	is.Equal(midtrans.Sandbox, midclient.APIEnvType)
}

func TestSnapCreateTokenQuick(t *testing.T) {
	is := is.New(t)
    now := time.Now()
    timestamp := strconv.FormatInt(now.Unix(), 10)

    midclient := midtrans.NewClient()
    midclient.ServerKey = "SB-Mid-server-GwUP_WGbJPXsDzsNEBRs8IYA"
    midclient.ClientKey = "SB-Mid-client-61XuGAwQ8Bj8LxSS"
    midclient.APIEnvType = midtrans.Sandbox

	var snapGateway midtrans.SnapGateway
	snapGateway = midtrans.SnapGateway{
		Client: midclient,
	}

	snapResp, err := snapGateway.GetTokenQuick("order-id-go-"+timestamp, 200000)
	var snapToken string
	if err != nil {
		snapToken = snapResp.Token
		is.OK(snapToken)
	}
}
