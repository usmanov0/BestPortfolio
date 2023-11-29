package blogService

import "BestPortfolio/internal/repo"

type BlogUseCase struct {
	blogRepo *repo.BlogRepo
}

func (b BlogUseCase) NewBlog(title, description, body string) {
	//TODO implement me
	panic("implement me")
}

func (b BlogUseCase) DeleteBlog(id uint64) error {
	//TODO implement me
	panic("implement me")
}

func NewBlogUseCase(blogRepo *repo.BlogRepo) BlogService {
	return &BlogUseCase{blogRepo: blogRepo}
}
