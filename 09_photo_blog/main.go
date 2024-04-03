package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public",
		http.FileServer(http.Dir("./public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func getCookie(w http.ResponseWriter, req *http.Request) http.Cookie {
	cookie, err := req.Cookie("session")
	cookie.MaxAge = -1

	if err == http.ErrNoCookie {
		sId := uuid.NewString()

		cookie = &http.Cookie{
			Name:  "session",
			Value: sId,
		}

		http.SetCookie(w, cookie)
	}

	return *cookie
}

func checkAndSetImage(cookie http.Cookie, fname string, w http.ResponseWriter, req *http.Request) http.Cookie {
	if !strings.Contains(cookie.Value, fname) {
		cookie.Value += "|" + fname
	}

	http.SetCookie(w, &cookie)

	return cookie
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)

	if req.Method == http.MethodPost {
		mf, fh, err := req.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close()

		// Create sha for filename
		ext := strings.Split(fh.Filename, ".")[1]
		h := sha1.New()
		io.Copy(h, mf)
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

		// Create new file
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		path := filepath.Join(wd, "public", "pics", fname)
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()

		// copy
		mf.Seek(0, 0)
		io.Copy(nf, mf)
		c = checkAndSetImage(c, fname, w, req)
	}

	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}
