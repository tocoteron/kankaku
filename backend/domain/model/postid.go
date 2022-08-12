package model

import "fmt"

type PostID struct {
	value string
}

func NewPostID(id string) (*PostID, error) {
	if err := validatePostID(id); err != nil {
		return nil, err
	}

	pid := &PostID{
		value: id,
	}

	return pid, nil
}

func (pid PostID) Equals(other PostID) bool {
	return pid.value == other.value
}

func (pid PostID) String() string {
	return pid.value
}

func validatePostID(id string) error {
	if len(id) == 0 {
		return fmt.Errorf("id must be not empty")
	}

	if len(id) > 256 {
		return fmt.Errorf("id must be 256 charactes or less")
	}

	return nil
}
