package memstorage

import (
	usersDomain "temp-prj/internal/domain/users"
	"temp-prj/internal/repository/errors"
)

type Storage struct {
	users map[string]usersDomain.User
}

func New() *Storage {
	return &Storage{
		users: make(map[string]usersDomain.User),
	}
}

func (s *Storage) SaveUser(user usersDomain.User) error {
	for _, u := range s.users {
		if u.Email == user.Email {
			return errors.ErrUserAlreadyExists
		}
	}

	s.users[user.UID] = user
	return nil
}

func (s *Storage) GetUserByEmail(email string) (usersDomain.User, error) {
	for _, u := range s.users {
		if u.Email == email {
			return u, nil
		}
	}

	return usersDomain.User{}, errors.ErrUserNotFound
}

func (s *Storage) GetUserByUID(uid string) (usersDomain.User, error) {
	user, ok := s.users[uid]

	if !ok {
		return usersDomain.User{}, errors.ErrUserNotFound
	}

	return user, nil
}
