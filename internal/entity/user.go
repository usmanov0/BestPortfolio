package entity

type User struct {
	Id         int
	Address    string
	LikedBlogs []*Blog
}
