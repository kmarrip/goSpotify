package mocks

import "github.com/stretchr/testify/mock"

type SpotifyMock struct {
	mock.Mock
}

func (m SpotifyMock) CurrentSong(accessToken string) (string, error) {
	args := m.Called(accessToken)
	return args.String(0), args.Error(1)
}

func (m SpotifyMock) Profile(accessToken string) (string, error) {
	args := m.Called(accessToken)
	return args.String(0), args.Error(1)
}
