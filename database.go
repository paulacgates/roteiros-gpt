package main

import (
	"database/sql"
	"embed"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"

	ora "github.com/sijms/go-ora/v2"
)

var (
	DB  *sql.DB
	SQL map[string]string

	//go:embed "sql/*"
	embeddedSQL embed.FS
)

const tracelogFile = "/var/log/trace.log"

func criaDB(usuarioOra string, senhaOra string, ipOra string, portaOra int, serviceOra string) {
	connectionString := criaConnectionString(usuarioOra, senhaOra, ipOra, portaOra, serviceOra)
	var err error
	DB, err = sql.Open("oracle", connectionString)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalln("ERROR: Teste de conexão não passou.", err)
	}

	log.Println("Conexão com", serviceOra, "OK")

	carregaArquivosSQL()
}

func criaConnectionString(usuarioOra string, senhaOra string, ipOra string, portaOra int, serviceOra string) string {
	var urlOptions map[string]string
	if runtime.GOOS == "windows" {
		removeTraceLog()
		urlOptions = map[string]string{
			"language":   "BRAZILIAN PORTUGUESE",
			"territory":  "BRAZIL",
			"LOB FETCH":  "POST",
			"TRACE FILE": tracelogFile,
		}
		url := ora.BuildUrl(ipOra, portaOra, serviceOra, usuarioOra, senhaOra, urlOptions)
		return url
	} else {
		urlOptions = map[string]string{
			"language":  "BRAZILIAN PORTUGUESE",
			"territory": "BRAZIL",
			"LOB FETCH": "POST",
		}
	}
	url := ora.BuildUrl(ipOra, portaOra, serviceOra, usuarioOra, senhaOra, urlOptions)
	return url
}
func removeTraceLog() {
	os.Remove(tracelogFile)
}

func carregaArquivosSQL() {
	SQL = make(map[string]string)
	matches, err := fs.Glob(embeddedSQL, "sql/*.sql")
	if err != nil {
		log.Fatalln("ERRO na leitura do diretório sql.", err)
	}

	for idx := range matches {
		bytes, err := fs.ReadFile(embeddedSQL, matches[idx])
		if err != nil {
			log.Fatalln("ERRO na leitura do arquivo SQL.", matches[idx], err)
		}

		name := filepath.Base(matches[idx])
		name = name[:len(name)-4]
		SQL[name] = string(bytes)
	}
	log.Println("Carregamento dos arquivos SQL OK")
}
