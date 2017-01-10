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
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    t.once.Do(func() {
        t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
    })

    data := map[string]interface{}{
        "Host": r.Host,
    }

    t.templ.Execute(w, data)
}