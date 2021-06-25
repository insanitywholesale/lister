package frontend

import (
	"context"
	pb "gitlab.com/insanitywholesale/lister/proto/v1"
	"google.golang.org/grpc"
	"html/template"
	"log"
	"net/http"
	"strings"
	"os"
	"path/filepath"
)

var (
	templatePath string
	lc           pb.ListerClient
)

func init() {
	templatePath = os.Getenv("TEMPLATE_PATH")
	if templatePath == "" {
		templatePath = "./frontend/templates"
	}
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
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		return
	}
	if r.Method == "POST" {
		var title string
		var items []string
		r.ParseForm()
		var itemString string
		if (r.Form["title"] != nil) && (len(r.Form["title"][0]) > 0) {
			title = r.Form["title"][0]
		} else {
			http.Error(w, "Please set a title for the list", 400)
		}
		if r.Form["items"] != nil {
			itemString = r.Form["items"][0]
			items = strings.Split(itemString, ",")
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
	mainPath := filepath.Join(templatePath, "main.html")
	listPath := filepath.Join(templatePath, "list.html")

	l, err := lc.GetAllLists(context.Background(), &pb.Empty{})
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	t, err := template.ParseFiles(mainPath, listPath)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	t.ExecuteTemplate(w, "main", l)
	return
}

func ShowForm(w http.ResponseWriter, r *http.Request) {
	mainPath := filepath.Join(templatePath, "main.html")
	formPath := filepath.Join(templatePath, "form.html")

	t, err := template.ParseFiles(mainPath, formPath)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	t.ExecuteTemplate(w, "main", nil)
	return
}
