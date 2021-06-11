package connexion

import (
	"net/http"

	render "./renders"
)

func Profil() {
	http.HandleFunc("/user_posts.html", render.Render_User_Posts)
}
