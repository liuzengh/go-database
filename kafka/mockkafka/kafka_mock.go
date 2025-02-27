// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mockkafka is a generated GoMock package.
package mockkafka

import (
	context "context"
	sarama "github.com/Shopify/sarama"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Produce mocks base method
func (m *MockClient) Produce(ctx context.Context, key, value []byte, headers ...sarama.RecordHeader) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key, value}
	for _, a := range headers {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Produce", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Produce indicates an expected call of Produce
func (mr *MockClientMockRecorder) Produce(ctx, key, value interface{}, headers ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key, value}, headers...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockClient)(nil).Produce), varargs...)
}

// SendMessage mocks base method
func (m *MockClient) SendMessage(ctx context.Context, topic string, key, value []byte, headers ...sarama.RecordHeader) (int32, int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, topic, key, value}
	for _, a := range headers {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendMessage", varargs...)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SendMessage indicates an expected call of SendMessage
func (mr *MockClientMockRecorder) SendMessage(ctx, topic, key, value interface{}, headers ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, topic, key, value}, headers...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockClient)(nil).SendMessage), varargs...)
}

// AsyncSendMessage mocks base method
func (m *MockClient) AsyncSendMessage(ctx context.Context, topic string, key, value []byte, headers ...sarama.RecordHeader) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, topic, key, value}
	for _, a := range headers {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AsyncSendMessage", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AsyncSendMessage indicates an expected call of AsyncSendMessage
func (mr *MockClientMockRecorder) AsyncSendMessage(ctx, topic, key, value interface{}, headers ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, topic, key, value}, headers...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AsyncSendMessage", reflect.TypeOf((*MockClient)(nil).AsyncSendMessage), varargs...)
}

// SendSaramaMessage mocks base method
func (m *MockClient) SendSaramaMessage(ctx context.Context, sMsg sarama.ProducerMessage) (int32, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendSaramaMessage", ctx, sMsg)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SendSaramaMessage indicates an expected call of SendSaramaMessage
func (mr *MockClientMockRecorder) SendSaramaMessage(ctx, sMsg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSaramaMessage", reflect.TypeOf((*MockClient)(nil).SendSaramaMessage), ctx, sMsg)
}
