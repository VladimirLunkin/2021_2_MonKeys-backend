// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	models "dripapp/internal/dripapp/models"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// AddMatch provides a mock function with given fields: ctx, firstUser, secondUser
func (_m *UserRepository) AddMatch(ctx context.Context, firstUser uint64, secondUser uint64) error {
	ret := _m.Called(ctx, firstUser, secondUser)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64) error); ok {
		r0 = rf(ctx, firstUser, secondUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddReaction provides a mock function with given fields: ctx, currentUserId, swipedUserId, reactionType
func (_m *UserRepository) AddReaction(ctx context.Context, currentUserId uint64, swipedUserId uint64, reactionType uint64) error {
	ret := _m.Called(ctx, currentUserId, swipedUserId, reactionType)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64, uint64) error); ok {
		r0 = rf(ctx, currentUserId, swipedUserId, reactionType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddReport provides a mock function with given fields: ctx, report
func (_m *UserRepository) AddReport(ctx context.Context, report models.NewReport) error {
	ret := _m.Called(ctx, report)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.NewReport) error); ok {
		r0 = rf(ctx, report)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUser provides a mock function with given fields: ctx, logUserData
func (_m *UserRepository) CreateUser(ctx context.Context, logUserData models.LoginUser) (models.User, error) {
	ret := _m.Called(ctx, logUserData)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(context.Context, models.LoginUser) models.User); ok {
		r0 = rf(ctx, logUserData)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.LoginUser) error); ok {
		r1 = rf(ctx, logUserData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMatches provides a mock function with given fields: ctx, firstUser, secondUser
func (_m *UserRepository) DeleteMatches(ctx context.Context, firstUser uint64, secondUser uint64) error {
	ret := _m.Called(ctx, firstUser, secondUser)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64) error); ok {
		r0 = rf(ctx, firstUser, secondUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteReaction provides a mock function with given fields: ctx, firstUser, secondUser
func (_m *UserRepository) DeleteReaction(ctx context.Context, firstUser uint64, secondUser uint64) error {
	ret := _m.Called(ctx, firstUser, secondUser)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, uint64) error); ok {
		r0 = rf(ctx, firstUser, secondUser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetLikes provides a mock function with given fields: ctx, currentUserId
func (_m *UserRepository) GetLikes(ctx context.Context, currentUserId uint64) ([]uint64, error) {
	ret := _m.Called(ctx, currentUserId)

	var r0 []uint64
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []uint64); ok {
		r0 = rf(ctx, currentUserId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, currentUserId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNextUserForSwipe provides a mock function with given fields: ctx, currentUser
func (_m *UserRepository) GetNextUserForSwipe(ctx context.Context, currentUser models.User) ([]models.User, error) {
	ret := _m.Called(ctx, currentUser)

	var r0 []models.User
	if rf, ok := ret.Get(0).(func(context.Context, models.User) []models.User); ok {
		r0 = rf(ctx, currentUser)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.User) error); ok {
		r1 = rf(ctx, currentUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReportDesc provides a mock function with given fields: ctx, reportId
func (_m *UserRepository) GetReportDesc(ctx context.Context, reportId uint64) (string, error) {
	ret := _m.Called(ctx, reportId)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, uint64) string); ok {
		r0 = rf(ctx, reportId)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, reportId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReports provides a mock function with given fields: ctx
func (_m *UserRepository) GetReports(ctx context.Context) (map[uint64]string, error) {
	ret := _m.Called(ctx)

	var r0 map[uint64]string
	if rf, ok := ret.Get(0).(func(context.Context) map[uint64]string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[uint64]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReportsCount provides a mock function with given fields: ctx, userId
func (_m *UserRepository) GetReportsCount(ctx context.Context, userId uint64) (uint64, error) {
	ret := _m.Called(ctx, userId)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, uint64) uint64); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReportsWithMaxCount provides a mock function with given fields: ctx, userId
func (_m *UserRepository) GetReportsWithMaxCount(ctx context.Context, userId uint64) (uint64, error) {
	ret := _m.Called(ctx, userId)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, uint64) uint64); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTags provides a mock function with given fields: ctx
func (_m *UserRepository) GetTags(ctx context.Context) (map[uint64]string, error) {
	ret := _m.Called(ctx)

	var r0 map[uint64]string
	if rf, ok := ret.Get(0).(func(context.Context) map[uint64]string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[uint64]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: ctx, email
func (_m *UserRepository) GetUser(ctx context.Context, email string) (models.User, error) {
	ret := _m.Called(ctx, email)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(context.Context, string) models.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: ctx, userID
func (_m *UserRepository) GetUserByID(ctx context.Context, userID uint64) (models.User, error) {
	ret := _m.Called(ctx, userID)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(context.Context, uint64) models.User); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsersLikes provides a mock function with given fields: ctx, currentUserId
func (_m *UserRepository) GetUsersLikes(ctx context.Context, currentUserId uint64) ([]models.User, error) {
	ret := _m.Called(ctx, currentUserId)

	var r0 []models.User
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []models.User); ok {
		r0 = rf(ctx, currentUserId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, currentUserId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsersMatches provides a mock function with given fields: ctx, currentUserId
func (_m *UserRepository) GetUsersMatches(ctx context.Context, currentUserId uint64) ([]models.User, error) {
	ret := _m.Called(ctx, currentUserId)

	var r0 []models.User
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []models.User); ok {
		r0 = rf(ctx, currentUserId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, currentUserId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsersMatchesWithSearching provides a mock function with given fields: ctx, currentUserId, searchTmpl
func (_m *UserRepository) GetUsersMatchesWithSearching(ctx context.Context, currentUserId uint64, searchTmpl string) ([]models.User, error) {
	ret := _m.Called(ctx, currentUserId, searchTmpl)

	var r0 []models.User
	if rf, ok := ret.Get(0).(func(context.Context, uint64, string) []models.User); ok {
		r0 = rf(ctx, currentUserId, searchTmpl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64, string) error); ok {
		r1 = rf(ctx, currentUserId, searchTmpl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateImgs provides a mock function with given fields: ctx, id, imgs
func (_m *UserRepository) UpdateImgs(ctx context.Context, id uint64, imgs []string) error {
	ret := _m.Called(ctx, id, imgs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, []string) error); ok {
		r0 = rf(ctx, id, imgs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateReportStatus provides a mock function with given fields: ctx, userId, reportStatus
func (_m *UserRepository) UpdateReportStatus(ctx context.Context, userId uint64, reportStatus string) error {
	ret := _m.Called(ctx, userId, reportStatus)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, string) error); ok {
		r0 = rf(ctx, userId, reportStatus)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: ctx, newUserData
func (_m *UserRepository) UpdateUser(ctx context.Context, newUserData models.User) (models.User, error) {
	ret := _m.Called(ctx, newUserData)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(context.Context, models.User) models.User); ok {
		r0 = rf(ctx, newUserData)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.User) error); ok {
		r1 = rf(ctx, newUserData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
