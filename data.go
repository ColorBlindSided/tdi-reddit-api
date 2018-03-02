package main

type Page struct {
	Data struct {
		Children Posts `json:"children"`
	}
}

type Post struct {
	Data Post_Data `json:"data"`
}

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

type Posts []Post

type Post_Collection struct {
	Matches Posts `json:"matches"`
}