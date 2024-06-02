package main

import (
    "html/template"
    "net/http"
    "sync"
)

var (
    tpl   = template.Must(template.ParseFiles("index.html"))
    tasks = struct {
        sync.Mutex
        List []string
    }{}
)

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/add", addHandler)
    http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    tasks.Lock()
    defer tasks.Unlock()
    tpl.Execute(w, tasks.List)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        task := r.FormValue("task")
        if task != "" {
            tasks.Lock()
            tasks.List = append(tasks.List, task)
            tasks.Unlock()
        }
    }
    http.Redirect(w, r, "/", http.StatusSeeOther)
}
