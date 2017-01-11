package main

import (
    "sync"
    "html/template"
    "net/http"
    "path/filepath"
)

type templateHandler struct {
    once sync.Once
    filename string
    templ *template.Template
    data map[string]interface{}
    dataInitializer func(t *templateHandler)
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    t.once.Do(func() {
        t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
    })

    if t.dataInitializer != nil {
        t.dataInitializer(t)
    } else {
        t.data = make(map[string]interface{})
    }

    t.data["Host"] = r.Host

    t.templ.Execute(w, t.data)
}