package secret

import (
	"encore.dev/rlog"
	"secret-min/lib"
)

// var secrets struct {
// 	SomeSecret    string
// 	AnotherSecret string
// }

// encore:service
type Service struct {
}

func initService() (*Service, error) {
	// rlog.Info("seceret values loaded", "secret1", secrets.SomeSecret, "secret2", secrets.AnotherSecret)
	l := lib.GetSecrets()
	rlog.Info("from lib", "data", l)

	return &Service{}, nil
}
