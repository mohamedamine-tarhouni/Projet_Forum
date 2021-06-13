package render

//the info needed to connect the user to other data
type USER struct {
	ID        int
	User_name string
}

//the columns from the table "Commentaire"
type Commentaire struct {
	ID_Com int
	User   USER
	Date   string
	Text   string
}

//the columns from the table "Post"
type Post struct {
	ID_Post     int
	User        USER
	Title       string
	Category    string
	Description string
	Img         string
	Comments    []Commentaire
}

//the Data needed to display the good posts with proper user and check the Status of the user
type Data struct {
	Posts    []Post
	Category string
	Status   string
}

//the Errors that a user can encounter while he signs up
type Errors struct {
	Err_name      string
	Err_surname   string
	Err_User_name string
	Err_Email     string
	Err_password  string
	Err_Cpassword string
}
