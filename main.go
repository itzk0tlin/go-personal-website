package main

/*
	Main file with routes and nothing else.
*/

import (
	"log"
	"net/http"
	"os"

	"github.com/dixxe/personal-website/iternal/controllers"
	custommiddleware "github.com/dixxe/personal-website/iternal/customMiddleware"
	"github.com/dixxe/personal-website/iternal/customMiddleware/cron"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/crypto/acme/autocert"
)

var domains = os.Args[1:]

func main() {

	log.Println(`
━━━━━━━━━━━━━━━━━━┏━┓━━━━┏┓━┏┓━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┏┓━━━━━━━━━━━━━┏┓━━━━━━━┏┓━━━━━
━━━━━━━━━━━━━━━━━━┃┏┛━━━━┃┃━┃┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┏┛┗┓━━━━━━━━━━━━┃┃━━━━━━┏┛┗┓━━━━
┏━┓━┏━━┓┏┓┏┓┏━━┓━┏┛┗┓┏━━┓┃┃━┗┛┏━━┓━━━━┏━━┓┏━━┓━┏━━┓┏━━┓━━━━┗┓┏┛┏━━┓┏┓┏┓┏━━┓┃┃━┏━━┓━┗┓┏┛┏━━┓
┃┏┓┓┃┏┓┃┃┗┛┃┗━┓┃━┗┓┏┛┃┏┓┃┃┃━━━┃━━┫━━━━┃┏┓┃┗━┓┃━┃┏┓┃┃┏┓┃━━━━━┃┃━┃┏┓┃┃┗┛┃┃┏┓┃┃┃━┗━┓┃━━┃┃━┃┏┓┃
┃┃┃┃┃┃━┫┗┓┏┛┃┗┛┗┓━┃┃━┃┃━┫┃┗┓━━┣━━┃━━━━┃┗┛┃┃┗┛┗┓┃┗┛┃┃┃━┫━━━━━┃┗┓┃┃━┫┃┃┃┃┃┗┛┃┃┗┓┃┗┛┗┓━┃┗┓┃┃━┫
┗┛┗┛┗━━┛━┗┛━┗━━━┛━┗┛━┗━━┛┗━┛━━┗━━┛━━━━┃┏━┛┗━━━┛┗━┓┃┗━━┛━━━━━┗━┛┗━━┛┗┻┻┛┃┏━┛┗━┛┗━━━┛━┗━┛┗━━┛
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃┃━━━━━━━┏━┛┃━━━━━━━━━━━━━━━━━━━━┃┃━━━━━━━━━━━━━━━━━━
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┗┛━━━━━━━┗━━┛━━━━━━━━━━━━━━━━━━━━┗┛━━━━━━━━━━━━━━━━━━
	   	`)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.StripSlashes)
	r.Use(custommiddleware.Statistics)
	r.Use(custommiddleware.DDOSProtection)

	r.Get("/", controllers.GetIndexHandler)

	r.Get("/blog", controllers.GetShowBlog)
	r.Get("/post/{id}", controllers.GetPost)
	r.Post("/post", controllers.PostCreatePost)
	r.Post("/post/delete", controllers.PostDeletePost)
	r.Post("/post/update", controllers.PostUpdatePost)

	r.Post("/admin/login", controllers.PostAdminLogin)
	r.Get("/admin", controllers.GetAdminLogin)

	r.Post("/cctweaked", controllers.PostSendMessage)
	r.Get("/cctweaked", controllers.GetMessage)
	r.Post("/cctweaked-beta", controllers.PostSendMessage)
	r.Get("/cctweaked-beta", controllers.GetMessageBeta)

	fs := http.FileServer(http.Dir("web/static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	log.Println(domains)

	launchCronJobs()

	if domains[0] == "" {

		log.Println("HTTP website started")
		log.Fatal(http.ListenAndServe("localhost:8080", r))

	} else {

		log.Println("HTTPS website started with HTTP redirect.")
		go redirectHTTPServer()
		for _, domain := range domains {
			log.Fatal(http.Serve(autocert.NewListener(domain), r))
		}
	}
}

func redirectHTTPServer() {
	if err := http.ListenAndServe(":8080", http.HandlerFunc(redirectTLS)); err != nil {
		log.Fatalf("ListenAndServe error: %v", err)
	}
}

func redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+domains[0]+":443"+r.RequestURI, http.StatusMovedPermanently)
}

func launchCronJobs() {
	log.Println("Launching cron jobs...")

	go cron.CleanRecentConnections()
}
