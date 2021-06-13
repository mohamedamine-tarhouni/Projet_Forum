package render

import "database/sql"

//this function deletes a post using its ID
func Delete_Post_BY_ID(db *sql.DB, ID_Post int) *sql.DB {
	query := "DELETE FROM Post WHERE ID_Post=?"
	_, _ = db.Exec(query, ID_Post)
	return db
}
