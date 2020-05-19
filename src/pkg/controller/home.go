package controller

import (
	"context"
	"html/template"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gauravgahlot/tink-wizard/src/client"
	"github.com/gauravgahlot/tink-wizard/src/pkg"
	"github.com/gauravgahlot/tink-wizard/src/pkg/types"
)

const (
	errTemplateExecute = "failed to execute template"
)

type home struct {
	templates map[string]*template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/about", h.handleAbout)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	tmpls, err := client.ListTemplates(context.Background())
	if err != nil {
		log.Error(err)
	}

	for _, tmp := range tmpls {
		log.Info(tmp.LastUpdated)
	}

	data := types.Home{}
	data.Title = "Home"
	err = h.templates[index].Execute(w, data)
	pkg.CheckError(err, errTemplateExecute)
}

func (h home) handleAbout(w http.ResponseWriter, r *http.Request) {
	data := types.Home{}
	data.Title = "About"
	err := h.templates[about].Execute(w, data)
	pkg.CheckError(err, errTemplateExecute)
}
