package main

import (
	"fmt"
	"net/http"

	"gopkg.in/ini.v1"
)

type configurationListHandler struct {
	appMode        string
	privateKeyPath string
	publicKeyPath  string
	paypalURL      string
	paypalCertPath string
	dbCon          string
	dbUser         string
	dbPassword     string
}

func (h *configurationListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h2>I am a GO application running inside Kubernetes.<h2> <h3>My properties are:</h3><ul>")
	fmt.Fprintf(w, "<li>app_mode: "+h.appMode+"</li>")
	fmt.Fprintf(w, "<li>private_key: "+h.privateKeyPath+"</li>")
	fmt.Fprintf(w, "<li>public_key: "+h.publicKeyPath+"</li>")
	fmt.Fprintf(w, "<li>paypal_url: "+h.paypalURL+"</li>")
	fmt.Fprintf(w, "<li>paypal_cert: "+h.paypalCertPath+"</li>")
	fmt.Fprintf(w, "<li>db_con: "+h.dbCon+"</li>")
	fmt.Fprintf(w, "<li>db_user: "+h.dbUser+"</li>")
	fmt.Fprintf(w, "<li>db_password: "+h.dbPassword+"</li>")
	fmt.Fprintf(w, "</ol>")

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK\n")

}

func main() {

	cfg, err := ini.LooseLoad("/config/settings.ini", "settings.ini")
	if err != nil {
		fmt.Printf("Failed to read configuration file: %v", err)
	}

	fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())

	clh := configurationListHandler{}
	clh.appMode = cfg.Section("").Key("app_mode").String()
	clh.privateKeyPath = cfg.Section("security").Key("private_key").String()
	clh.publicKeyPath = cfg.Section("security").Key("public_key").String()
	clh.paypalURL = cfg.Section("paypal").Key("paypal_url").String()
	clh.paypalCertPath = cfg.Section("paypal").Key("paypal_cert").String()
	clh.dbCon = cfg.Section("mysql").Key("db_con").String()
	clh.dbUser = cfg.Section("mysql").Key("db_user").String()
	clh.dbPassword = cfg.Section("mysql").Key("db_password").String()

	fmt.Println("Simple web server is starting on port 8080...")

	http.Handle("/", &clh)
	http.HandleFunc("/health", healthHandler)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server at port 8080: %v", err)
	}
}
