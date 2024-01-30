package middlewares

import (
	"net/http"
	"path/filepath"
	"strings"
	"text/template"
)

var Templates *template.Template

func InitHTMLRendering(templateDir string) {
	Templates = template.Must(template.ParseGlob(templateDir + "/*.html"))
}

func HTMLRenderingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".html") {
			err := Templates.ExecuteTemplate(w, filepath.Base(r.URL.Path), nil)
			if err != nil {
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				return
			}
			return
		}
		next.ServeHTTP(w, r)
	})
}
