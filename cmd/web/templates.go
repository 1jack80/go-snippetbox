package main 

import "snippetbox.1jack80/internal/models"

type templateData struct {
    Snippet *models.Snippet
    Snippets []*models.Snippet
}
