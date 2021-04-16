// Code generated by MockGen. DO NOT EDIT.
// Source: server.go

// Package mock_controllers is a generated GoMock package.
package mock_controllers

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	"github.com/gofrs/uuid"
	distributedpb "px.dev/pixie/src/carnot/planner/distributedpb"
	plannerpb "px.dev/pixie/src/carnot/planner/plannerpb"
	planpb "px.dev/pixie/src/carnot/planpb"
	queryresultspb "px.dev/pixie/src/carnot/queryresultspb"
	carnotpb "px.dev/pixie/src/carnot/carnotpb"
	querybrokerpb "px.dev/pixie/src/vizier/services/query_broker/querybrokerpb"
	tracker "px.dev/pixie/src/vizier/services/query_broker/tracker"
)

// MockPlanner is a mock of Planner interface.
type MockPlanner struct {
	ctrl     *gomock.Controller
	recorder *MockPlannerMockRecorder
}

// MockPlannerMockRecorder is the mock recorder for MockPlanner.
type MockPlannerMockRecorder struct {
	mock *MockPlanner
}

// NewMockPlanner creates a new mock instance.
func NewMockPlanner(ctrl *gomock.Controller) *MockPlanner {
	mock := &MockPlanner{ctrl: ctrl}
	mock.recorder = &MockPlannerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlanner) EXPECT() *MockPlannerMockRecorder {
	return m.recorder
}

// CompileMutations mocks base method.
func (m *MockPlanner) CompileMutations(planState *distributedpb.LogicalPlannerState, request *plannerpb.CompileMutationsRequest) (*plannerpb.CompileMutationsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompileMutations", planState, request)
	ret0, _ := ret[0].(*plannerpb.CompileMutationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompileMutations indicates an expected call of CompileMutations.
func (mr *MockPlannerMockRecorder) CompileMutations(planState, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompileMutations", reflect.TypeOf((*MockPlanner)(nil).CompileMutations), planState, request)
}

// Free mocks base method.
func (m *MockPlanner) Free() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Free")
}

// Free indicates an expected call of Free.
func (mr *MockPlannerMockRecorder) Free() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Free", reflect.TypeOf((*MockPlanner)(nil).Free))
}

// Plan mocks base method.
func (m *MockPlanner) Plan(planState *distributedpb.LogicalPlannerState, req *plannerpb.QueryRequest) (*distributedpb.LogicalPlannerResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Plan", planState, req)
	ret0, _ := ret[0].(*distributedpb.LogicalPlannerResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Plan indicates an expected call of Plan.
func (mr *MockPlannerMockRecorder) Plan(planState, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Plan", reflect.TypeOf((*MockPlanner)(nil).Plan), planState, req)
}

// MockExecutor is a mock of Executor interface.
type MockExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockExecutorMockRecorder
}

// MockExecutorMockRecorder is the mock recorder for MockExecutor.
type MockExecutorMockRecorder struct {
	mock *MockExecutor
}

// NewMockExecutor creates a new mock instance.
func NewMockExecutor(ctrl *gomock.Controller) *MockExecutor {
	mock := &MockExecutor{ctrl: ctrl}
	mock.recorder = &MockExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecutor) EXPECT() *MockExecutorMockRecorder {
	return m.recorder
}

// AddResult mocks base method.
func (m *MockExecutor) AddResult(res *querybrokerpb.AgentQueryResultRequest) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddResult", res)
}

// AddResult indicates an expected call of AddResult.
func (mr *MockExecutorMockRecorder) AddResult(res interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddResult", reflect.TypeOf((*MockExecutor)(nil).AddResult), res)
}

// AddStreamedResult mocks base method.
func (m *MockExecutor) AddStreamedResult(res *carnotpb.TransferResultChunkRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddStreamedResult", res)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddStreamedResult indicates an expected call of AddStreamedResult.
func (mr *MockExecutorMockRecorder) AddStreamedResult(res interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddStreamedResult", reflect.TypeOf((*MockExecutor)(nil).AddStreamedResult), res)
}

// ExecuteQuery mocks base method.
func (m *MockExecutor) ExecuteQuery(planMap map[uuid.UUID]*planpb.Plan, analyze bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteQuery", planMap, analyze)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecuteQuery indicates an expected call of ExecuteQuery.
func (mr *MockExecutorMockRecorder) ExecuteQuery(planMap, analyze interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteQuery", reflect.TypeOf((*MockExecutor)(nil).ExecuteQuery), planMap, analyze)
}

// GetQueryID mocks base method.
func (m *MockExecutor) GetQueryID() uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueryID")
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// GetQueryID indicates an expected call of GetQueryID.
func (mr *MockExecutorMockRecorder) GetQueryID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryID", reflect.TypeOf((*MockExecutor)(nil).GetQueryID))
}

// WaitForCompletion mocks base method.
func (m *MockExecutor) WaitForCompletion() (*queryresultspb.QueryResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForCompletion")
	ret0, _ := ret[0].(*queryresultspb.QueryResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForCompletion indicates an expected call of WaitForCompletion.
func (mr *MockExecutorMockRecorder) WaitForCompletion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForCompletion", reflect.TypeOf((*MockExecutor)(nil).WaitForCompletion))
}

// MockAgentsTracker is a mock of AgentsTracker interface.
type MockAgentsTracker struct {
	ctrl     *gomock.Controller
	recorder *MockAgentsTrackerMockRecorder
}

// MockAgentsTrackerMockRecorder is the mock recorder for MockAgentsTracker.
type MockAgentsTrackerMockRecorder struct {
	mock *MockAgentsTracker
}

// NewMockAgentsTracker creates a new mock instance.
func NewMockAgentsTracker(ctrl *gomock.Controller) *MockAgentsTracker {
	mock := &MockAgentsTracker{ctrl: ctrl}
	mock.recorder = &MockAgentsTrackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgentsTracker) EXPECT() *MockAgentsTrackerMockRecorder {
	return m.recorder
}

// GetAgentInfo mocks base method.
func (m *MockAgentsTracker) GetAgentInfo() tracker.AgentsInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAgentInfo")
	ret0, _ := ret[0].(tracker.AgentsInfo)
	return ret0
}

// GetAgentInfo indicates an expected call of GetAgentInfo.
func (mr *MockAgentsTrackerMockRecorder) GetAgentInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAgentInfo", reflect.TypeOf((*MockAgentsTracker)(nil).GetAgentInfo))
}
