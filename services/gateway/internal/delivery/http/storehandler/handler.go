package storehandler

import (
	"gatewayapp/internal/client/authclient"
	"gatewayapp/internal/client/storeclient"
)

type Handler struct {
	storeClient storeclient.Client
	authClient  authclient.Client
}

func New(storeClient storeclient.Client,
	authClient authclient.Client) Handler {
	return Handler{
		storeClient: storeClient,
		authClient:  authClient,
	}
}
