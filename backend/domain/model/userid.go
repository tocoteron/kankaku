package model

type UserID struct {
	value string
}

func NewUserID(id string) UserID {
	return UserID{
		value: id,
	}
}

func (uid UserID) Equals(other UserID) bool {
	return uid.value == other.value
}

func (uid UserID) String() string {
	return uid.value
}
