package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the applications config settings
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
