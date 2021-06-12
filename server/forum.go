package connexion

import (
	"net/http"

	render "./renders"
)

func Forum() {
	http.HandleFunc("/Posts.html", render.Render_Categories)
	http.HandleFunc("/informatics.html", render.Render_Posts)
	http.HandleFunc("/Culture.html", render.Render_Posts)
	http.HandleFunc("/Video Games.html", render.Render_Posts)
	http.HandleFunc("/Geography.html", render.Render_Posts)
	http.HandleFunc("/Music.html", render.Render_Posts)
	http.HandleFunc("/Post_informatics.html", render.Render_posting)
	http.HandleFunc("/Post_Culture.html", render.Render_posting)
	http.HandleFunc("/Post_Video Games.html", render.Render_posting)
	http.HandleFunc("/Post_Geography.html", render.Render_posting)
	http.HandleFunc("/Post_Music.html", render.Render_posting)
	http.HandleFunc("/comment_informatics.html", render.Render_commenting)
	http.HandleFunc("/comment_Culture.html", render.Render_commenting)
	http.HandleFunc("/comment_Video Games.html", render.Render_commenting)
	http.HandleFunc("/comment_Geography.html", render.Render_commenting)
	http.HandleFunc("/comment_Music.html", render.Render_commenting)
	http.HandleFunc("/informatics", render.Render_Upload)
	http.HandleFunc("/Culture", render.Render_Upload)
	http.HandleFunc("/Video Games", render.Render_Upload)
	http.HandleFunc("/Music", render.Render_Upload)
	http.HandleFunc("/Geography", render.Render_Upload)
}
