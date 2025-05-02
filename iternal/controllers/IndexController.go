/*
Controller that only shows index template! Simple as it is!
*/

package controllers

import (
	"context"
	"net/http"

	"github.com/dixxe/personal-website/web/templates"
)

// Controller that shows "/" root page.
func GetIndexHandler(w http.ResponseWriter, r *http.Request) {
	// Templates are components and this is basic way to render them.
	// I use this way across all code
	component := templates.IndexPage()
	component.Render(context.Background(), w)
}

func GetContactHandler(w http.ResponseWriter, r *http.Request) {
	// Templates are components and this is basic way to render them.
	// I use this way across all code
	component := templates.Contacts()
	component.Render(context.Background(), w)
}

func GetAboutHandler(w http.ResponseWriter, r *http.Request) {
	// Templates are components and this is basic way to render them.
	// I use this way across all code
	component := templates.MoreInfo()
	component.Render(context.Background(), w)
}
