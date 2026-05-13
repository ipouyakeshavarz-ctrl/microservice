package orderhandler

import (
	"gatewayapp/internal/client/authclient"
	"gatewayapp/internal/client/orderclient"
)

type Handler struct {
	orderClient orderclient.Client
	authClient  authclient.Client
}

func New(orderClient orderclient.Client,
	authClient authclient.Client) Handler {
	return Handler{
		orderClient: orderClient,
		authClient:  authClient,
	}
}
