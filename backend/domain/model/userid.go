package model

import "fmt"

type UserID struct {
	value string
}

func NewUserID(id string) (*UserID, error) {
	if err := validateUserID(id); err != nil {
		return nil, err
	}

	uid := &UserID{
		value: id,
	}

	return uid, nil
}

func (uid UserID) Equals(other UserID) bool {
	return uid.value == other.value
}

func (uid UserID) String() string {
	return uid.value
}

func validateUserID(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("id must be not empty")
	}

	if len(id) > 256 {
		return fmt.Errorf("id must be 256 charactes or less")
	}

	return nil
}
