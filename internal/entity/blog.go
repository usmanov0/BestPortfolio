package entity

type Blog struct {
	Id            int
	Title         string
	Description   string
	Body          string
	LikeCount     int
	ViewCount     int
	Date          string
	UnixTime      string
	LikedUsers    []*User
	IsLikedByUser bool `json:"-"`
}

func (blog *Blog) View() {
	blog.ViewCount++
}

func (blog *Blog) Like(user *User) {
	blog.LikedUsers = append(blog.LikedUsers, user)
	blog.LikeCount = len(blog.LikedUsers)
}

func (blog *Blog) CheckIsLikedByUser(user *User) {
	for _, likedUser := range blog.LikedUsers {
		if likedUser.Id == user.Id {
			blog.IsLikedByUser = true
			return
		}
	}
	blog.IsLikedByUser = false
}
