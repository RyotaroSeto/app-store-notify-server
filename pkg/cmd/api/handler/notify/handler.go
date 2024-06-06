package notify

import (
	"app-store-notify-server/pkg/pb/api"
	"app-store-notify-server/pkg/pb/api/apiconnect"
	"context"

	"app-store-notify-server/pkg/cmd/api/usecase/notify"

	"connectrpc.com/connect"
)

type handler struct {
	notifyUsecase notify.Usecase
}

func NewHandler(notifyUsecase notify.Usecase) apiconnect.NotifyServiceHandler {
	return &handler{
		notifyUsecase: notifyUsecase,
	}
}

func (h *handler) Notify(_ context.Context, req *connect.Request[api.NotifyRequest]) (*connect.Response[api.NotifyResponse], error) {
	if err := h.notifyUsecase.Notify(req.Msg.SignedPayload); err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}

	return connect.NewResponse(&api.NotifyResponse{}), nil
}
