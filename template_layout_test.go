package belajargolangweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Template Layout Golang Web

// Template Layout ParseFile Golang Web

func TemplateLayoutParseFile(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/layout.gohtml",
		"./templates/footer.gohtml",
	))

	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Title": "Tamplate Layout",
		"Name":  "Eko",
	})
}

func TestTemplateLayoutParseFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayoutParseFile(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// Template Layout ParseGlob Golang Web

func TemplateLayoutParseGlob(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))

	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Title": "Tamplate Layout",
		"Name":  "Eko",
	})
}

func TestTemplateLayoutParseGlob(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayoutParseGlob(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// Template Layout Embed Golang Web

func TemplateLayoutEmbed(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))

	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Title": "Tamplate Layout",
		"Name":  "Eko",
	})
}

func TestTemplateLayoutEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayoutEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
