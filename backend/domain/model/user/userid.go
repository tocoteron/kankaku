package user

type UserID struct {
	id string
}

func NewUserID(id string) UserID {
	return UserID{
		id: id,
	}
}

func (uid *UserID) Equals(other UserID) bool {
	return uid.id == other.id
}
