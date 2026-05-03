package producthandler

import (
	"gatewayapp/internal/client/authclient"
	"gatewayapp/internal/client/productclient"
)

type Handler struct {
	productClient productclient.Client
	authClient    authclient.Client
}

func New(productClient productclient.Client,
	authClient authclient.Client) Handler {
	return Handler{
		productClient: productClient,
		authClient:    authClient,
	}
}
