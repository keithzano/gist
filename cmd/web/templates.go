package main

import "gist/internal/models"

type templateData struct {
	Gist  *models.Gist
	Gists []*models.Gist
}
