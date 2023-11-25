package models

type Post struct {
	ID          int    `db:"id" validate:"-"`
	Title       string `db:"title" validate:"required,min=3,max=128"`
	Description string `db:"description" validate:"required,max=1024"`
	UserID      int    `db:"user_id" validate:"required"`
}

type Comment struct {
	ID        int    `db:"id" validate:"-"`
	Text      string `db:"text" validate:"required,max=1024"`
	UserID    int    `db:"user_id" validate:"required"`
	UserEmail string `db:"user_email" validate:"required,email,max=128"`
	PostID    int    `db:"post_id" validate:"required"`
}

type Like struct {
	ID        int `db:"id" validate:"-"`
	UserID    int `db:"user_id" validate:"required"`
	CommentID int `db:"comment_id" validate:"required"`
}
