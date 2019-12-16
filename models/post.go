package models

type Post struct {
	Id int
	Title string
	Content string
	CreatedAt string
}

var posts = []Post{
	Post{Id: 1, Title: "Test title 1", Content: "Test content 12345", CreatedAt: "12 december"},
	Post{Id: 2, Title: "Test title 2", Content: "Test content 12345", CreatedAt: "13 december"},
	Post{Id: 3, Title: "Test title 3", Content: "Test content 12345", CreatedAt: "14 december"},
	Post{Id: 4, Title: "Test title 4", Content: "Test content 12345", CreatedAt: "15 december"},
	Post{Id: 5, Title: "Test title 5", Content: "Test content 12345", CreatedAt: "16 december"},
}

func GetPosts() []Post {
	return posts
}