package carthandler

import (
	"gatewayapp/internal/client/authclient"
	"gatewayapp/internal/client/cartclient"
)

type Handler struct {
	cartClient cartclient.Client
	authClient authclient.Client
}

func New(cartClient cartclient.Client,
	authClient authclient.Client) Handler {
	return Handler{
		cartClient: cartClient,
		authClient: authClient,
	}
}
