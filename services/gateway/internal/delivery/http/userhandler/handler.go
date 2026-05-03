package userhandler

import (
	"gatewayapp/internal/client/authclient"
	"gatewayapp/internal/client/userclient"
)

type Handler struct {
	userClient userclient.Client
	authClient authclient.Client
}

func New(userClient userclient.Client, authClient authclient.Client) Handler {
	return Handler{
		userClient: userClient,
		authClient: authClient,
	}
}
