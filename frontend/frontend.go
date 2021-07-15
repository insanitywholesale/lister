package frontend

import (
	"context"
	"embed"
	pb "gitlab.com/insanitywholesale/lister/proto/v1"
	"google.golang.org/grpc"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	//go:embed templates
	templates    embed.FS
	templatefs   fs.FS
	lc           pb.ListerClient
)

func init() {
	listerName := os.Getenv("LISTER_NAME")
	if listerName == "" {
		listerName = "localhost"
	}
	listerPort := os.Getenv("LISTER_PORT")
	if listerPort == "" {
		listerPort = "15200"
	}
	listerAddr := listerName + ":" + listerPort

	conn, err := grpc.Dial(listerAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	lc = pb.NewListerClient(conn)
	tfs, err := fs.Sub(templates, "templates")
	if err != nil {
		log.Fatal(err)
	}
	templatefs = tfs
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("this a form submission endpoint, wtf u doin fam\n"))
		return
	}
	if r.Method == "POST" {
		var title string
		var items []string
		var separator string
		r.ParseForm()
		var itemString string
		if (r.Form["title"] != nil) && (len(r.Form["title"][0]) > 0) {
			title = r.Form["title"][0]
		} else {
			http.Error(w, "Please set a title for the list", 400)
		}
		if r.Form["separator"] != nil {
			separator = r.Form["separator"][0]
		} else {
			separator = " "
		}
		if r.Form["items"] != nil {
			itemString = r.Form["items"][0]
			items = strings.Split(itemString, separator)
		} else {
			http.Error(w, "Please add some items to the list", 400)
		}
		log.Println(title, items)
		_, err := lc.AddList(context.Background(), &pb.List{Title: title, Items: items})
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		http.Redirect(w, r, "/ui", 301)
		return
	}
	return
}

func ShowLists(w http.ResponseWriter, r *http.Request) {
	l, err := lc.GetAllLists(context.Background(), &pb.Empty{})
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	tfs, err := template.ParseFS(templatefs, "main.html", "list.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	tfs.ExecuteTemplate(w, "main", l)
	return
}

func ShowForm(w http.ResponseWriter, r *http.Request) {
	tfs, err := template.ParseFS(templatefs, "main.html", "list.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	tfs.ExecuteTemplate(w, "main", nil)
	return
}
