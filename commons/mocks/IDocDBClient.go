// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	documentDB_client "github.eagleview.com/engineering/symphony-service/commons/documentDB_client"

	testing "testing"
)

// IDocDBClient is an autogenerated mock type for the IDocDBClient type
type IDocDBClient struct {
	mock.Mock
}

// BuildQueryForCallBack provides a mock function with given fields: ctx, event, status, workflowID, stepID, TaskName, callbackResponse
func (_m *IDocDBClient) BuildQueryForCallBack(ctx context.Context, event string, status string, workflowID string, stepID string, TaskName string, callbackResponse map[string]interface{}) (interface{}, interface{}) {
	ret := _m.Called(ctx, event, status, workflowID, stepID, TaskName, callbackResponse)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, string, map[string]interface{}) interface{}); ok {
		r0 = rf(ctx, event, status, workflowID, stepID, TaskName, callbackResponse)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 interface{}
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string, string, map[string]interface{}) interface{}); ok {
		r1 = rf(ctx, event, status, workflowID, stepID, TaskName, callbackResponse)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(interface{})
		}
	}

	return r0, r1
}

// BuildQueryForUpdateWorkflowDataCallout provides a mock function with given fields: ctx, TaskName, stepID, status, starttime, IsWaitTask
func (_m *IDocDBClient) BuildQueryForUpdateWorkflowDataCallout(ctx context.Context, TaskName string, stepID string, status string, starttime int64, IsWaitTask bool) interface{} {
	ret := _m.Called(ctx, TaskName, stepID, status, starttime, IsWaitTask)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, int64, bool) interface{}); ok {
		r0 = rf(ctx, TaskName, stepID, status, starttime, IsWaitTask)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// CheckConnection provides a mock function with given fields: ctx
func (_m *IDocDBClient) CheckConnection(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchStepExecutionData provides a mock function with given fields: ctx, StepId
func (_m *IDocDBClient) FetchStepExecutionData(ctx context.Context, StepId string) (documentDB_client.StepExecutionDataBody, error) {
	ret := _m.Called(ctx, StepId)

	var r0 documentDB_client.StepExecutionDataBody
	if rf, ok := ret.Get(0).(func(context.Context, string) documentDB_client.StepExecutionDataBody); ok {
		r0 = rf(ctx, StepId)
	} else {
		r0 = ret.Get(0).(documentDB_client.StepExecutionDataBody)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, StepId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchWorkflowExecutionData provides a mock function with given fields: ctx, workFlowId
func (_m *IDocDBClient) FetchWorkflowExecutionData(ctx context.Context, workFlowId string) (documentDB_client.WorkflowExecutionDataBody, error) {
	ret := _m.Called(ctx, workFlowId)

	var r0 documentDB_client.WorkflowExecutionDataBody
	if rf, ok := ret.Get(0).(func(context.Context, string) documentDB_client.WorkflowExecutionDataBody); ok {
		r0 = rf(ctx, workFlowId)
	} else {
		r0 = ret.Get(0).(documentDB_client.WorkflowExecutionDataBody)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, workFlowId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHipsterCountPerDay provides a mock function with given fields: ctx
func (_m *IDocDBClient) GetHipsterCountPerDay(ctx context.Context) (int64, error) {
	ret := _m.Called(ctx)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context) int64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertStepExecutionData provides a mock function with given fields: ctx, StepExecutionData
func (_m *IDocDBClient) InsertStepExecutionData(ctx context.Context, StepExecutionData documentDB_client.StepExecutionDataBody) error {
	ret := _m.Called(ctx, StepExecutionData)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, documentDB_client.StepExecutionDataBody) error); ok {
		r0 = rf(ctx, StepExecutionData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertWorkflowExecutionData provides a mock function with given fields: ctx, Data
func (_m *IDocDBClient) InsertWorkflowExecutionData(ctx context.Context, Data documentDB_client.WorkflowExecutionDataBody) error {
	ret := _m.Called(ctx, Data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, documentDB_client.WorkflowExecutionDataBody) error); ok {
		r0 = rf(ctx, Data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDocumentDB provides a mock function with given fields: ctx, query, update, collectionName
func (_m *IDocDBClient) UpdateDocumentDB(ctx context.Context, query interface{}, update interface{}, collectionName string) error {
	ret := _m.Called(ctx, query, update, collectionName)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, string) error); ok {
		r0 = rf(ctx, query, update, collectionName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIDocDBClient creates a new instance of IDocDBClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewIDocDBClient(t testing.TB) *IDocDBClient {
	mock := &IDocDBClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
