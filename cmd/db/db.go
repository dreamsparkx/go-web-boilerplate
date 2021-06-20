package db

import "github.com/dreamsparkx/go-web-boilerplate/internal/config"

func ConnectDB(app *config.AppConfig) {
	(*app).DB = config.NewDatabase(config.PostgresDriver)
	conn := config.PgConnString{
		Host:            app.DbCredentials.Host,
		User:            app.DbCredentials.User,
		Password:        app.DbCredentials.Password,
		Name:            app.DbCredentials.Name,
		SSLModeDisabled: app.DbCredentials.SSLModeDisabled,
	}
	(*app).DB.Connect(conn.BuildConnectionString())
}
