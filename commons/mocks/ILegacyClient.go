// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	legacy_client "github.eagleview.com/engineering/symphony-service/lambdas/legacyupdate/legacy_client"

	testing "testing"
)

// ILegacyClient is an autogenerated mock type for the ILegacyClient type
type ILegacyClient struct {
	mock.Mock
}

// UpdateReportStatus provides a mock function with given fields: ctx, req
func (_m *ILegacyClient) UpdateReportStatus(ctx context.Context, req *legacy_client.LegacyUpdateRequest) error {
	ret := _m.Called(ctx, req)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *legacy_client.LegacyUpdateRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewILegacyClient creates a new instance of ILegacyClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewILegacyClient(t testing.TB) *ILegacyClient {
	mock := &ILegacyClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}