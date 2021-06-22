package main

import (
	"context"
	pb "gitlab.com/insanitywholesale/lister/proto/v1"
	"google.golang.org/grpc"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	templatePath string
	lc           pb.ListerClient
)

func main() {
	templatePath = os.Getenv("TEMPLATE_PATH")
	if templatePath == "" {
		templatePath = "./templates"
	}
	frontendPort := os.Getenv("FRONTEND_PORT")
	if frontendPort == "" {
		frontendPort = "8080"
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
	defer conn.Close()
	lc = pb.NewListerClient(conn)
	log.Println("connected to lister over grpc")

	http.HandleFunc("/", showLists)

	log.Fatal(http.ListenAndServe(":"+frontendPort, nil))
}

func showLists(w http.ResponseWriter, r *http.Request) {
	mainPath := filepath.Join(templatePath, "main.html")
	listPath := filepath.Join(templatePath, "list.html")

	l, err := lc.GetAllLists(context.Background(), &pb.Empty{})
	if err != nil {
		log.Println(err)
	}

	t, err := template.ParseFiles(mainPath, listPath)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	t.ExecuteTemplate(w, "main", l)
	return
}
