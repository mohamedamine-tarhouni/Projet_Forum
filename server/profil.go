package connexion

import (
	"net/http"

	render "./renders"
)

//this function calls Handlefunc only to render the User_posts page
func Profil() {
	http.HandleFunc("/user_posts.html", render.Render_User_Posts)
}
