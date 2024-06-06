package appstore

import (
	iap_appstore "github.com/awa/go-iap/appstore"

	"app-store-notify-server/pkg/domain/appstore"
)

type verifier struct {
	client *iap_appstore.Client
}

func NewVerifier() appstore.Verifier {
	return &verifier{
		client: iap_appstore.New(),
	}
}

// https://github.com/awa/go-iapを利用することで署名を検証できる
func (v *verifier) ParseNotification(signedPayload string) (*appstore.Notification, error) {
	// decodedPayloadにはAppStoreからデコードされたデータが取得できる
	var decodedPayload iap_appstore.SubscriptionNotificationV2DecodedPayload
	if err := v.client.ParseNotificationV2WithClaim(signedPayload, &decodedPayload); err != nil {
		return nil, err
	}

	if err := decodedPayload.Valid(); err != nil {
		return nil, err
	}

	return &appstore.Notification{
		NotificationType: toNotificationType(decodedPayload),
	}, nil
}

func toNotificationType(decodedPayload iap_appstore.SubscriptionNotificationV2DecodedPayload) appstore.NotificationType {
	switch decodedPayload.NotificationType {
	case iap_appstore.NotificationTypeV2Test:
		return appstore.NotificationTypeTest
	default:
		return 0
	}
}
