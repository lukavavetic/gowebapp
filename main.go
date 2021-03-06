package main

import (
	"creatingWebApp/webapp/controller"
	"creatingWebApp/webapp/middleware"
	"creatingWebApp/webapp/model"
	"database/sql"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := connectToDb()
	defer db.Close()

	model.UpdateUserEmail("matija.vavetic@gmail.com")

	result, err := model.GetUserById()


	if err != nil {
		panic("wtf")
	}

	fmt.Println(result.Email)

	templates := populateTemplates()

	controller.StartUp(templates)

	http.ListenAndServeTLS(":8000","cert.pem","key.pem", &middleware.TimeoutMiddleware{Next: new(middleware.GzipMiddleware)})
}

func connectToDb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@/gowebapp")
	if err != nil {
		log.Fatalln(fmt.Errorf("unable to connecto to database: %v", err))
	}

	model.SetDatabase(db)

	return db
}


func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)

	const basePath = "templates"

	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))

	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))

	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}

	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}

	for _, fi := range fis {

		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}

		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}

		f.Close()

		tmpl := template.Must(layout.Clone())

		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}

		result[fi.Name()] = tmpl
	}

	return result
}
