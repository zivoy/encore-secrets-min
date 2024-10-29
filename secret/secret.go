package secret

import "encore.dev/rlog"

var secrets struct {
	SomeSecret    string
	AnotherSecret string
}

// encore:service
type Service struct {
}

func initService() (*Service, error) {
	rlog.Info("seceret values loaded", "secret1", secrets.SomeSecret, "secret2", secrets.AnotherSecret)

	return &Service{}, nil
}
