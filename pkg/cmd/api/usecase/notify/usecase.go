package notify

import (
	"app-store-notify-server/pkg/domain/appstore"
	"fmt"
	"log"
)

type Usecase interface {
	Notify(signedPayload string) error
}

type usecase struct {
	appstoreVerifier appstore.Verifier
}

func NewUsecase(appstoreVerifier appstore.Verifier) Usecase {
	return &usecase{
		appstoreVerifier: appstoreVerifier,
	}
}

func (u *usecase) Notify(signedPayload string) error {
	notification, err := u.appstoreVerifier.ParseNotification(signedPayload)
	if err != nil {
		return err
	}

	switch notification.NotificationType {
	case appstore.NotificationTypeTest:
		log.Println("parse notification succeeded! test notification received!")
	default:
		return fmt.Errorf(fmt.Sprintf("unknown notification type: %d", notification.NotificationType))
	}

	return nil
}
