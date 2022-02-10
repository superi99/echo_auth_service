package main

import (
	"github.com/labstack/echo/v4"
	"io"
	"skillspar/user_service/db"
	"skillspar/user_service/handler"
	"skillspar/user_service/server"
	"skillspar/user_service/store/impl"
	"text/template"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	sv := server.New()

	// session
	//var sessionStore = sessions.NewCookieStore([]byte(utils.Getenv("SESSION_SECRET", "SECRET")))
	//sessionStore.Options = &sessions.Options{
	//	Path:     "/",
	//	MaxAge:   86400 * 7,
	//	HttpOnly: true,
	//}
	//sv.Use(session.Middleware(sessionStore))

	// html template
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	sv.Renderer = t

	d := db.New()

	db.AutoMigrate(d)
	//db.Seeding(d)

	us := impl.NewUserStore(d)
	cs := impl.NewClientStore(d)

	h := handler.NewHandler(us, cs)

	routesV1 := sv.Group("")

	h.Register(routesV1)

	sv.Logger.Fatal(sv.Start("127.0.0.1:8000"))
}
