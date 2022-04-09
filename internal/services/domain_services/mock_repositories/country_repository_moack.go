package mock_repositories

import (
	"reflect"
	"reservation-api/internal/commons"
	"reservation-api/internal/dto"
	"reservation-api/internal/models"

	"github.com/golang/mock/gomock"
)

// MockCountryRepository is a mock of CountryRepository interface.
type MockCountryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCountryRepositoryMockRecorder
}

// MockCountryRepositoryMockRecorder is the mock recorder for MockCountryRepository.
type MockCountryRepositoryMockRecorder struct {
	mock *MockCountryRepository
}

// NewMockCountryRepository creates a new mock instance.
func NewMockCountryRepository(ctrl *gomock.Controller) *MockCountryRepository {
	mock := &MockCountryRepository{ctrl: ctrl}
	mock.recorder = &MockCountryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCountryRepository) EXPECT() *MockCountryRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCountryRepository) Create(arg0 *models.Country) (*models.Country, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*models.Country)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCountryRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCountryRepository)(nil).Create), arg0)
}

// Find mocks base method.
func (m *MockCountryRepository) Find(arg0 uint64) (*models.Country, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(*models.Country)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockCountryRepositoryMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCountryRepository)(nil).Find), arg0)
}

// FindAll mocks base method.
func (m *MockCountryRepository) FindAll(arg0 *dto.PaginationFilter) (*commons.PaginatedList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0)
	ret0, _ := ret[0].(*commons.PaginatedList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockCountryRepositoryMockRecorder) FindAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockCountryRepository)(nil).FindAll), arg0)
}

// GetProvinces mocks base method.
func (m *MockCountryRepository) GetProvinces(arg0 uint64) ([]*models.Province, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProvinces", arg0)
	ret0, _ := ret[0].([]*models.Province)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProvinces indicates an expected call of GetProvinces.
func (mr *MockCountryRepositoryMockRecorder) GetProvinces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProvinces", reflect.TypeOf((*MockCountryRepository)(nil).GetProvinces), arg0)
}

// Update mocks base method.
func (m *MockCountryRepository) Update(arg0 *models.Country) (*models.Country, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*models.Country)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCountryRepositoryMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCountryRepository)(nil).Update), arg0)
}
