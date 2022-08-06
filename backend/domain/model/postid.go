package model

type PostID struct {
	value string
}

func NewPostID(id string) PostID {
	return PostID{
		value: id,
	}
}

func (pid PostID) Equals(other PostID) bool {
	return pid.value == other.value
}

func (pid PostID) String() string {
	return pid.value
}
