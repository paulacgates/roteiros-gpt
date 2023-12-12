package main

import (
	"embed"
	"log"
	"os"
	"strconv"
	"text/template"

	"github.com/joho/godotenv"
)

//go:embed web/html/* web/static/*
var embeddedFiles embed.FS

var templates *template.Template

func main() {
	config := carregaConfig()
	log.Println("Iniciando...")
	criaDB(config.UsuarioOra, config.SenhaOra, config.IPOra, config.PortaOra, config.ServiceOra)
	carregaTemplates()
	log.Printf("Carregamento dos Templates OK")
	log.Printf("Ouvindo na porta %d\n", config.Porta)
	log.Println("Pronto!")
	//log.Fatalln(http.ListenAndServe(addr, roteador))
}

func carregaConfig() config {
	godotenv.Load(".env")
	strPortaOra := os.Getenv("PORTA_ORA")
	portaOra, err := strconv.Atoi(strPortaOra)
	if err != nil {
		log.Fatalln("PORTA_ORA inválida:", err)
	}
	strPorta := os.Getenv("PORTA")
	porta, err := strconv.Atoi(strPorta)
	if err != nil {
		log.Fatalln("PORTA inválida.", err)
	}
	return config{
		LogDest:    os.Getenv("LOG"),
		UsuarioOra: os.Getenv("USUARIO_ORA"),
		SenhaOra:   os.Getenv("SENHA_ORA"),
		IPOra:      os.Getenv("IP_ORA"),
		PortaOra:   portaOra,
		ServiceOra: os.Getenv("SERVICE_ORA"),
		Porta:      porta,
		CertPath:   "",
		KeyPath:    "",
	}
}

func carregaTemplates() {
	var err error
	templates, err = template.ParseFS(embeddedFiles, "web/html/*.html", "web/html/fragments/*.html")
	if err != nil {
		log.Fatal(err)
	}
}

type config struct {
	LogDest    string
	UsuarioOra string
	SenhaOra   string
	IPOra      string
	PortaOra   int
	ServiceOra string
	Porta      int
	CertPath   string
	KeyPath    string
}
