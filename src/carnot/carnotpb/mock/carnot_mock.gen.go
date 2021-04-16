// Code generated by MockGen. DO NOT EDIT.
// Source: carnot.pb.go

// Package mock_carnotpb is a generated GoMock package.
package mock_carnotpb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	carnotpb "px.dev/pixie/src/carnot/carnotpb"
)

// MockisTransferResultChunkRequest_Result is a mock of isTransferResultChunkRequest_Result interface.
type MockisTransferResultChunkRequest_Result struct {
	ctrl     *gomock.Controller
	recorder *MockisTransferResultChunkRequest_ResultMockRecorder
}

// MockisTransferResultChunkRequest_ResultMockRecorder is the mock recorder for MockisTransferResultChunkRequest_Result.
type MockisTransferResultChunkRequest_ResultMockRecorder struct {
	mock *MockisTransferResultChunkRequest_Result
}

// NewMockisTransferResultChunkRequest_Result creates a new mock instance.
func NewMockisTransferResultChunkRequest_Result(ctrl *gomock.Controller) *MockisTransferResultChunkRequest_Result {
	mock := &MockisTransferResultChunkRequest_Result{ctrl: ctrl}
	mock.recorder = &MockisTransferResultChunkRequest_ResultMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockisTransferResultChunkRequest_Result) EXPECT() *MockisTransferResultChunkRequest_ResultMockRecorder {
	return m.recorder
}

// Equal mocks base method.
func (m *MockisTransferResultChunkRequest_Result) Equal(arg0 interface{}) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockisTransferResultChunkRequest_ResultMockRecorder) Equal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockisTransferResultChunkRequest_Result)(nil).Equal), arg0)
}

// MarshalTo mocks base method.
func (m *MockisTransferResultChunkRequest_Result) MarshalTo(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalTo", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalTo indicates an expected call of MarshalTo.
func (mr *MockisTransferResultChunkRequest_ResultMockRecorder) MarshalTo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalTo", reflect.TypeOf((*MockisTransferResultChunkRequest_Result)(nil).MarshalTo), arg0)
}

// Size mocks base method.
func (m *MockisTransferResultChunkRequest_Result) Size() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size.
func (mr *MockisTransferResultChunkRequest_ResultMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockisTransferResultChunkRequest_Result)(nil).Size))
}

// isTransferResultChunkRequest_Result mocks base method.
func (m *MockisTransferResultChunkRequest_Result) isTransferResultChunkRequest_Result() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isTransferResultChunkRequest_Result")
}

// isTransferResultChunkRequest_Result indicates an expected call of isTransferResultChunkRequest_Result.
func (mr *MockisTransferResultChunkRequest_ResultMockRecorder) isTransferResultChunkRequest_Result() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isTransferResultChunkRequest_Result", reflect.TypeOf((*MockisTransferResultChunkRequest_Result)(nil).isTransferResultChunkRequest_Result))
}

// MockisTransferResultChunkRequest_SinkResult_ResultContents is a mock of isTransferResultChunkRequest_SinkResult_ResultContents interface.
type MockisTransferResultChunkRequest_SinkResult_ResultContents struct {
	ctrl     *gomock.Controller
	recorder *MockisTransferResultChunkRequest_SinkResult_ResultContentsMockRecorder
}

// MockisTransferResultChunkRequest_SinkResult_ResultContentsMockRecorder is the mock recorder for MockisTransferResultChunkRequest_SinkResult_ResultContents.
type MockisTransferResultChunkRequest_SinkResult_ResultContentsMockRecorder struct {
	mock *MockisTransferResultChunkRequest_SinkResult_ResultContents
}

// NewMockisTransferResultChunkRequest_SinkResult_ResultContents creates a new mock instance.
func NewMockisTransferResultChunkRequest_SinkResult_ResultContents(ctrl *gomock.Controller) *MockisTransferResultChunkRequest_SinkResult_ResultContents {
	mock := &MockisTransferResultChunkRequest_SinkResult_ResultContents{ctrl: ctrl}
	mock.recorder = &MockisTransferResultChunkRequest_SinkResult_ResultContentsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockisTransferResultChunkRequest_SinkResult_ResultContents) EXPECT() *MockisTransferResultChunkRequest_SinkResult_ResultContentsMockRecorder {
	return m.recorder
}

// Equal mocks base method.
func (m *MockisTransferResultChunkRequest_SinkResult_ResultContents) Equal(arg0 interface{}) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockisTransferResultChunkRequest_SinkResult_ResultContentsMockRecorder) Equal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockisTransferResultChunkRequest_SinkResult_ResultContents)(nil).Equal), arg0)
}

// MarshalTo mocks base method.
func (m *MockisTransferResultChunkRequest_SinkResult_ResultContents) MarshalTo(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalTo", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalTo indicates an expected call of MarshalTo.
func (mr *MockisTransferResultChunkRequest_SinkResult_ResultContentsMockRecorder) MarshalTo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalTo", reflect.TypeOf((*MockisTransferResultChunkRequest_SinkResult_ResultContents)(nil).MarshalTo), arg0)
}

// Size mocks base method.
func (m *MockisTransferResultChunkRequest_SinkResult_ResultContents) Size() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size.
func (mr *MockisTransferResultChunkRequest_SinkResult_ResultContentsMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockisTransferResultChunkRequest_SinkResult_ResultContents)(nil).Size))
}

// isTransferResultChunkRequest_SinkResult_ResultContents mocks base method.
func (m *MockisTransferResultChunkRequest_SinkResult_ResultContents) isTransferResultChunkRequest_SinkResult_ResultContents() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isTransferResultChunkRequest_SinkResult_ResultContents")
}

// isTransferResultChunkRequest_SinkResult_ResultContents indicates an expected call of isTransferResultChunkRequest_SinkResult_ResultContents.
func (mr *MockisTransferResultChunkRequest_SinkResult_ResultContentsMockRecorder) isTransferResultChunkRequest_SinkResult_ResultContents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isTransferResultChunkRequest_SinkResult_ResultContents", reflect.TypeOf((*MockisTransferResultChunkRequest_SinkResult_ResultContents)(nil).isTransferResultChunkRequest_SinkResult_ResultContents))
}

// MockisTransferResultChunkRequest_SinkResult_Destination is a mock of isTransferResultChunkRequest_SinkResult_Destination interface.
type MockisTransferResultChunkRequest_SinkResult_Destination struct {
	ctrl     *gomock.Controller
	recorder *MockisTransferResultChunkRequest_SinkResult_DestinationMockRecorder
}

// MockisTransferResultChunkRequest_SinkResult_DestinationMockRecorder is the mock recorder for MockisTransferResultChunkRequest_SinkResult_Destination.
type MockisTransferResultChunkRequest_SinkResult_DestinationMockRecorder struct {
	mock *MockisTransferResultChunkRequest_SinkResult_Destination
}

// NewMockisTransferResultChunkRequest_SinkResult_Destination creates a new mock instance.
func NewMockisTransferResultChunkRequest_SinkResult_Destination(ctrl *gomock.Controller) *MockisTransferResultChunkRequest_SinkResult_Destination {
	mock := &MockisTransferResultChunkRequest_SinkResult_Destination{ctrl: ctrl}
	mock.recorder = &MockisTransferResultChunkRequest_SinkResult_DestinationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockisTransferResultChunkRequest_SinkResult_Destination) EXPECT() *MockisTransferResultChunkRequest_SinkResult_DestinationMockRecorder {
	return m.recorder
}

// Equal mocks base method.
func (m *MockisTransferResultChunkRequest_SinkResult_Destination) Equal(arg0 interface{}) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockisTransferResultChunkRequest_SinkResult_DestinationMockRecorder) Equal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockisTransferResultChunkRequest_SinkResult_Destination)(nil).Equal), arg0)
}

// MarshalTo mocks base method.
func (m *MockisTransferResultChunkRequest_SinkResult_Destination) MarshalTo(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarshalTo", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalTo indicates an expected call of MarshalTo.
func (mr *MockisTransferResultChunkRequest_SinkResult_DestinationMockRecorder) MarshalTo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalTo", reflect.TypeOf((*MockisTransferResultChunkRequest_SinkResult_Destination)(nil).MarshalTo), arg0)
}

// Size mocks base method.
func (m *MockisTransferResultChunkRequest_SinkResult_Destination) Size() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size.
func (mr *MockisTransferResultChunkRequest_SinkResult_DestinationMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockisTransferResultChunkRequest_SinkResult_Destination)(nil).Size))
}

// isTransferResultChunkRequest_SinkResult_Destination mocks base method.
func (m *MockisTransferResultChunkRequest_SinkResult_Destination) isTransferResultChunkRequest_SinkResult_Destination() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isTransferResultChunkRequest_SinkResult_Destination")
}

// isTransferResultChunkRequest_SinkResult_Destination indicates an expected call of isTransferResultChunkRequest_SinkResult_Destination.
func (mr *MockisTransferResultChunkRequest_SinkResult_DestinationMockRecorder) isTransferResultChunkRequest_SinkResult_Destination() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isTransferResultChunkRequest_SinkResult_Destination", reflect.TypeOf((*MockisTransferResultChunkRequest_SinkResult_Destination)(nil).isTransferResultChunkRequest_SinkResult_Destination))
}

// MockResultSinkServiceClient is a mock of ResultSinkServiceClient interface.
type MockResultSinkServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockResultSinkServiceClientMockRecorder
}

// MockResultSinkServiceClientMockRecorder is the mock recorder for MockResultSinkServiceClient.
type MockResultSinkServiceClientMockRecorder struct {
	mock *MockResultSinkServiceClient
}

// NewMockResultSinkServiceClient creates a new mock instance.
func NewMockResultSinkServiceClient(ctrl *gomock.Controller) *MockResultSinkServiceClient {
	mock := &MockResultSinkServiceClient{ctrl: ctrl}
	mock.recorder = &MockResultSinkServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResultSinkServiceClient) EXPECT() *MockResultSinkServiceClientMockRecorder {
	return m.recorder
}

// TransferResultChunk mocks base method.
func (m *MockResultSinkServiceClient) TransferResultChunk(ctx context.Context, opts ...grpc.CallOption) (carnotpb.ResultSinkService_TransferResultChunkClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TransferResultChunk", varargs...)
	ret0, _ := ret[0].(carnotpb.ResultSinkService_TransferResultChunkClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferResultChunk indicates an expected call of TransferResultChunk.
func (mr *MockResultSinkServiceClientMockRecorder) TransferResultChunk(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferResultChunk", reflect.TypeOf((*MockResultSinkServiceClient)(nil).TransferResultChunk), varargs...)
}

// MockResultSinkService_TransferResultChunkClient is a mock of ResultSinkService_TransferResultChunkClient interface.
type MockResultSinkService_TransferResultChunkClient struct {
	ctrl     *gomock.Controller
	recorder *MockResultSinkService_TransferResultChunkClientMockRecorder
}

// MockResultSinkService_TransferResultChunkClientMockRecorder is the mock recorder for MockResultSinkService_TransferResultChunkClient.
type MockResultSinkService_TransferResultChunkClientMockRecorder struct {
	mock *MockResultSinkService_TransferResultChunkClient
}

// NewMockResultSinkService_TransferResultChunkClient creates a new mock instance.
func NewMockResultSinkService_TransferResultChunkClient(ctrl *gomock.Controller) *MockResultSinkService_TransferResultChunkClient {
	mock := &MockResultSinkService_TransferResultChunkClient{ctrl: ctrl}
	mock.recorder = &MockResultSinkService_TransferResultChunkClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResultSinkService_TransferResultChunkClient) EXPECT() *MockResultSinkService_TransferResultChunkClientMockRecorder {
	return m.recorder
}

// CloseAndRecv mocks base method.
func (m *MockResultSinkService_TransferResultChunkClient) CloseAndRecv() (*carnotpb.TransferResultChunkResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseAndRecv")
	ret0, _ := ret[0].(*carnotpb.TransferResultChunkResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseAndRecv indicates an expected call of CloseAndRecv.
func (mr *MockResultSinkService_TransferResultChunkClientMockRecorder) CloseAndRecv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseAndRecv", reflect.TypeOf((*MockResultSinkService_TransferResultChunkClient)(nil).CloseAndRecv))
}

// CloseSend mocks base method.
func (m *MockResultSinkService_TransferResultChunkClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockResultSinkService_TransferResultChunkClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockResultSinkService_TransferResultChunkClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockResultSinkService_TransferResultChunkClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockResultSinkService_TransferResultChunkClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockResultSinkService_TransferResultChunkClient)(nil).Context))
}

// Header mocks base method.
func (m *MockResultSinkService_TransferResultChunkClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockResultSinkService_TransferResultChunkClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockResultSinkService_TransferResultChunkClient)(nil).Header))
}

// RecvMsg mocks base method.
func (m_2 *MockResultSinkService_TransferResultChunkClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockResultSinkService_TransferResultChunkClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockResultSinkService_TransferResultChunkClient)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockResultSinkService_TransferResultChunkClient) Send(arg0 *carnotpb.TransferResultChunkRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockResultSinkService_TransferResultChunkClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockResultSinkService_TransferResultChunkClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockResultSinkService_TransferResultChunkClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockResultSinkService_TransferResultChunkClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockResultSinkService_TransferResultChunkClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockResultSinkService_TransferResultChunkClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockResultSinkService_TransferResultChunkClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockResultSinkService_TransferResultChunkClient)(nil).Trailer))
}

// MockResultSinkServiceServer is a mock of ResultSinkServiceServer interface.
type MockResultSinkServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockResultSinkServiceServerMockRecorder
}

// MockResultSinkServiceServerMockRecorder is the mock recorder for MockResultSinkServiceServer.
type MockResultSinkServiceServerMockRecorder struct {
	mock *MockResultSinkServiceServer
}

// NewMockResultSinkServiceServer creates a new mock instance.
func NewMockResultSinkServiceServer(ctrl *gomock.Controller) *MockResultSinkServiceServer {
	mock := &MockResultSinkServiceServer{ctrl: ctrl}
	mock.recorder = &MockResultSinkServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResultSinkServiceServer) EXPECT() *MockResultSinkServiceServerMockRecorder {
	return m.recorder
}

// TransferResultChunk mocks base method.
func (m *MockResultSinkServiceServer) TransferResultChunk(arg0 carnotpb.ResultSinkService_TransferResultChunkServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferResultChunk", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// TransferResultChunk indicates an expected call of TransferResultChunk.
func (mr *MockResultSinkServiceServerMockRecorder) TransferResultChunk(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferResultChunk", reflect.TypeOf((*MockResultSinkServiceServer)(nil).TransferResultChunk), arg0)
}

// MockResultSinkService_TransferResultChunkServer is a mock of ResultSinkService_TransferResultChunkServer interface.
type MockResultSinkService_TransferResultChunkServer struct {
	ctrl     *gomock.Controller
	recorder *MockResultSinkService_TransferResultChunkServerMockRecorder
}

// MockResultSinkService_TransferResultChunkServerMockRecorder is the mock recorder for MockResultSinkService_TransferResultChunkServer.
type MockResultSinkService_TransferResultChunkServerMockRecorder struct {
	mock *MockResultSinkService_TransferResultChunkServer
}

// NewMockResultSinkService_TransferResultChunkServer creates a new mock instance.
func NewMockResultSinkService_TransferResultChunkServer(ctrl *gomock.Controller) *MockResultSinkService_TransferResultChunkServer {
	mock := &MockResultSinkService_TransferResultChunkServer{ctrl: ctrl}
	mock.recorder = &MockResultSinkService_TransferResultChunkServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResultSinkService_TransferResultChunkServer) EXPECT() *MockResultSinkService_TransferResultChunkServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockResultSinkService_TransferResultChunkServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockResultSinkService_TransferResultChunkServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockResultSinkService_TransferResultChunkServer)(nil).Context))
}

// Recv mocks base method.
func (m *MockResultSinkService_TransferResultChunkServer) Recv() (*carnotpb.TransferResultChunkRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*carnotpb.TransferResultChunkRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockResultSinkService_TransferResultChunkServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockResultSinkService_TransferResultChunkServer)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockResultSinkService_TransferResultChunkServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockResultSinkService_TransferResultChunkServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockResultSinkService_TransferResultChunkServer)(nil).RecvMsg), m)
}

// SendAndClose mocks base method.
func (m *MockResultSinkService_TransferResultChunkServer) SendAndClose(arg0 *carnotpb.TransferResultChunkResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAndClose", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAndClose indicates an expected call of SendAndClose.
func (mr *MockResultSinkService_TransferResultChunkServerMockRecorder) SendAndClose(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAndClose", reflect.TypeOf((*MockResultSinkService_TransferResultChunkServer)(nil).SendAndClose), arg0)
}

// SendHeader mocks base method.
func (m *MockResultSinkService_TransferResultChunkServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockResultSinkService_TransferResultChunkServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockResultSinkService_TransferResultChunkServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockResultSinkService_TransferResultChunkServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockResultSinkService_TransferResultChunkServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockResultSinkService_TransferResultChunkServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockResultSinkService_TransferResultChunkServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockResultSinkService_TransferResultChunkServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockResultSinkService_TransferResultChunkServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockResultSinkService_TransferResultChunkServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockResultSinkService_TransferResultChunkServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockResultSinkService_TransferResultChunkServer)(nil).SetTrailer), arg0)
}
