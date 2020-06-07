package registration

import (
	"database/sql"

	registration_v1 "github.com/motapratik/golang-rest-api-demo/pkg/api/registration/v1"
	"github.com/motapratik/golang-rest-api-demo/pkg/http_multiplexer"
)

// AllRoutes of registration
func AllRoutes(
	multiplexer http_multiplexer.IMultiplexer,
	dbConnection *sql.DB,
) {
	registrationService := registration_v1.NewService(dbConnection)
	registration_v1.Routes(multiplexer, registrationService)
}
