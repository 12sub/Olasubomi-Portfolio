package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func main() {
	var err error
	tmpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}

	// Static files (images, css, js)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/", handleAbout)
	http.HandleFunc("/about", handleAbout)
	http.HandleFunc("/software", handleSoftware)
	http.HandleFunc("/projects", handleProjects)
	http.HandleFunc("/experience", handleExperience)
	http.HandleFunc("/contact", handleContact)

	log.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type PageData struct {
	Active string
	Body   template.HTML
}

func render(w http.ResponseWriter, r *http.Request, page string) {
	var buf bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buf, page, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// HTMX request → return only the content partial
	if r.Header.Get("HX-Request") == "true" {
		w.Write(buf.Bytes())
		return
	}

	// Full page load → wrap in base template
	data := PageData{Active: page, Body: template.HTML(buf.String())}
	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleAbout(w http.ResponseWriter, r *http.Request)       { render(w, r, "about") }
func handleSoftware(w http.ResponseWriter, r *http.Request)    { render(w, r, "software") }
func handleProjects(w http.ResponseWriter, r *http.Request)    { render(w, r, "projects") }
func handleExperience(w http.ResponseWriter, r *http.Request)  { render(w, r, "experience") }

func handleContact(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// In production, wire this to email/SMTP/Discord/etc.
		w.Write([]byte(`<div class="ok">✅ Message sent successfully! I will get back to you soon.</div>`))
		return
	}
	render(w, r, "contact")
}
