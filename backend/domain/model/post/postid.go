package post

type PostID struct {
	id string
}

func NewPostID(id string) PostID {
	return PostID{
		id: id,
	}
}

func (pid *PostID) Equals(other PostID) bool {
	return pid.id == other.id
}
