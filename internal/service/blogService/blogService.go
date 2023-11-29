package blogService

type BlogService interface {
	NewBlog(title, description, body string)
	DeleteBlog(id uint64) error
}
