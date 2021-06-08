package connexion

import (
	"net/http"

	render "./renders"
)

func Forum() {
	// http.HandleFunc("/Post.html", RenderTemplate_accueil)
	http.HandleFunc("/informatique.html", render.Render_Posts)
	http.HandleFunc("/Post_informatique.html", render.Render_posting)
	http.HandleFunc("/comment.html", render.Render_commenting)
}
