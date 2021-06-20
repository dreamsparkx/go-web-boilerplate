package session

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/dreamsparkx/go-web-boilerplate/internal/config"
)

func CreateSession(app *config.AppConfig) *scs.SessionManager {
	var session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	return session
}
