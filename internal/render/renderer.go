package render

import (
    "html/template"
    "log"
    "net/http"
    "path/filepath"

    "github.com/gin-gonic/gin"
)

var templates = make(map[string]*template.Template)

func LoadTemplates() {
	layoutUser := "web/templates/layouts/layout_user.html"
	layoutAdmin := "web/templates/layouts/layout_admin.html"

	// Load user templates
	userPages, _ := filepath.Glob("web/templates/pages/user/*.html")
	for _, page := range userPages {
		name := filepath.Base(page)
		tmpl, err := template.ParseFiles(layoutUser, page)
		if err != nil {
			log.Printf("❌ Lỗi load user template %s: %v\n", page, err)
			continue
		}
		templates["user_"+name] = tmpl
	}

	// Load admin templates
	adminPages, _ := filepath.Glob("web/templates/pages/admin/*.html")
	for _, page := range adminPages {
		name := filepath.Base(page)
		tmpl, err := template.ParseFiles(layoutAdmin, page)
		if err != nil {
			log.Printf("❌ Lỗi load admin template %s: %v\n", page, err)
			continue
		}
		templates["admin_"+name] = tmpl
	}

	log.Println("✅ Templates loaded")
}

func RenderUser(c *gin.Context, tmpl string, data gin.H) {
    if t, ok := templates["user_"+tmpl+".html"]; ok {
        c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
        if err := t.Execute(c.Writer, data); err != nil {
            log.Println("❌ Lỗi render user template:", err)
            c.String(http.StatusInternalServerError, "Lỗi server")
        }
    } else {
        log.Println("❌ Template không tồn tại:", tmpl)
        c.String(http.StatusNotFound, "Trang không tồn tại")
    }
}

func RenderAdmin(c *gin.Context, tmpl string, data gin.H) {
    if t, ok := templates["admin_"+tmpl+".html"]; ok {
        c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
        if err := t.Execute(c.Writer, data); err != nil {
            log.Println("❌ Lỗi render admin template:", err)
            c.String(http.StatusInternalServerError, "Lỗi server")
        }
    } else {
        log.Println("❌ Template không tồn tại:", tmpl)
        c.String(http.StatusNotFound, "Trang không tồn tại")
    }
}