package main

//Designed to match the reddit JSON api format
type Page struct {
	Data struct {
		Children Posts `json:"children"`
	}
}

//Designed to match an individual post format in the api
type Post struct {
	Data Post_Data `json:"data"`
}

//The relevant data from a post, organized in order of importance
type Post_Data struct {
	Title			string	`json:"title"`
	Author 			string 	`json:"author"`
	Created			float64	`json:"created_utc"`
	Id				string	`json:"id"`
	Post_Hint		string	`json:"post_hint"`
	Selftext		string	`json:"selftext"`
	Domain			string 	`json:"domain"`
	URL				string 	`json:"url"`
	Thumbnail		string	`json:"thumbnail"`
	Score			int	`json:"score"`
	Comments		int	`json:"num_comments"`
	Crossposts		int	`json:"num_crossposts"`
	Gilded			int	`json:"gilded"`
	Spoiler			bool	`json:"spoiler"`
	Distinguished	string	`json:"distinguished"`
	Stickied		bool	`json:"stickied"`
	Archived		bool	`json:"archived"`
	Locked			bool	`json:"locked"`
}

//A useful type synonym
type Posts []Post

//A collection of posts. Used for displaying the matching results in JSON format.
type Post_Collection struct {
	Matches Posts `json:"matches"`
}