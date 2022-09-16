package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/beherasantosh/bookings/pkg/config"
	"github.com/beherasantosh/bookings/pkg/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string,  td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//get requested template frm cache
	t, ok := tc[tmpl]
	if !ok{
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)

	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(w)

	if err != nil{
		log.Println("error writing template to browser",err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error){
	mycache := make(map[string]*template.Template)

	//get all files with page.html
	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil{
		return mycache, err
	}

	for _, page := range pages{
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil{
			return mycache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil{
			return mycache, err
		}

		if len(matches)>0{
			ts, err  = ts.ParseGlob("./templates/*.layout.html")
			if err != nil{
				return mycache, err
			}
		}

		mycache[name] = ts

	}

	return mycache, nil
}



// RenderTemplate render templates using html template
// func RenderTemplate1(w http.ResponseWriter, tmpl string) {
// 	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.html")
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("error parsing template", err)
// 		return
// 	}
// }

// var tc = make(map[string]*template.Template)

// func RenderTemplateTest(w http.ResponseWriter, t string){
// 	var tmpl *template.Template
// 	var err error

// 	//check to see if we already have cache

// 	_, inMap := tc[t]
// 	if !inMap {
// 		//need to create a teamplte
// 		log.Println("creating template and adding to cache")
// 		err = createTemplateCache2(t)
// 		if err!= nil{
// 			log.Println(err)
// 		}

// 	}else{
// 		log.Println("using cached teamplete")
// 	}

// 	tmpl = tc[t]

	
// 	err = tmpl.Execute(w, nil)

// 	if err != nil {
// 		log.Println( err)
// 	}

// }

// func createTemplateCache2(t string) error{
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t), 
// 		"./templates/base.layout.html",
// 	}

// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	tc[t] = tmpl

// 	return nil
// }