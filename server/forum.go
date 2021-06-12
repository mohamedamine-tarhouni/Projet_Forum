package connexion

import (
	"net/http"

	render "./renders"
)

func Forum() {
	http.HandleFunc("/Posts.html", render.Render_Categories)
	http.HandleFunc("/informatique.html", render.Render_Posts)
	http.HandleFunc("/culture.html", render.Render_Posts)
	http.HandleFunc("/jeux_videos.html", render.Render_Posts)
	http.HandleFunc("/geography.html", render.Render_Posts)
	http.HandleFunc("/music.html", render.Render_Posts)
	http.HandleFunc("/Post_informatique.html", render.Render_posting)
	http.HandleFunc("/Post_culture.html", render.Render_posting)
	http.HandleFunc("/Post_jeux_videos.html", render.Render_posting)
	http.HandleFunc("/Post_geography.html", render.Render_posting)
	http.HandleFunc("/Post_music.html", render.Render_posting)
	http.HandleFunc("/comment_informatique.html", render.Render_commenting)
	http.HandleFunc("/comment_culture.html", render.Render_commenting)
	http.HandleFunc("/comment_jeux_videos.html", render.Render_commenting)
	http.HandleFunc("/comment_geography.html", render.Render_commenting)
	http.HandleFunc("/comment_music.html", render.Render_commenting)
	http.HandleFunc("/informatique", render.Render_Upload)
	http.HandleFunc("/culture", render.Render_Upload)
	http.HandleFunc("/jeux_videos", render.Render_Upload)
	http.HandleFunc("/music", render.Render_Upload)
	http.HandleFunc("/geography", render.Render_Upload)
}
