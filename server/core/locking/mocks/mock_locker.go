// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/core/locking (interfaces: Locker)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	locking "github.com/runatlantis/atlantis/server/core/locking"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockLocker struct {
	fail func(message string, callerSkip ...int)
}

func NewMockLocker(options ...pegomock.Option) *MockLocker {
	mock := &MockLocker{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockLocker) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockLocker) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockLocker) GetLock(_param0 string) (*models.ProjectLock, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockLocker().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetLock", params, []reflect.Type{reflect.TypeOf((**models.ProjectLock)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *models.ProjectLock
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*models.ProjectLock)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockLocker) List() (map[string]models.ProjectLock, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockLocker().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("List", params, []reflect.Type{reflect.TypeOf((*map[string]models.ProjectLock)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 map[string]models.ProjectLock
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(map[string]models.ProjectLock)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockLocker) TryLock(_param0 models.Project, _param1 string, _param2 models.PullRequest, _param3 models.User) (locking.TryLockResponse, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockLocker().")
	}
	params := []pegomock.Param{_param0, _param1, _param2, _param3}
	result := pegomock.GetGenericMockFrom(mock).Invoke("TryLock", params, []reflect.Type{reflect.TypeOf((*locking.TryLockResponse)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 locking.TryLockResponse
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(locking.TryLockResponse)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockLocker) Unlock(_param0 string) (*models.ProjectLock, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockLocker().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Unlock", params, []reflect.Type{reflect.TypeOf((**models.ProjectLock)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *models.ProjectLock
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*models.ProjectLock)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockLocker) UnlockByPull(_param0 string, _param1 int) ([]models.ProjectLock, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockLocker().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UnlockByPull", params, []reflect.Type{reflect.TypeOf((*[]models.ProjectLock)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []models.ProjectLock
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]models.ProjectLock)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockLocker) VerifyWasCalledOnce() *VerifierMockLocker {
	return &VerifierMockLocker{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockLocker) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockLocker {
	return &VerifierMockLocker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockLocker) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockLocker {
	return &VerifierMockLocker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockLocker) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockLocker {
	return &VerifierMockLocker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockLocker struct {
	mock                   *MockLocker
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockLocker) GetLock(_param0 string) *MockLocker_GetLock_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetLock", params, verifier.timeout)
	return &MockLocker_GetLock_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockLocker_GetLock_OngoingVerification struct {
	mock              *MockLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockLocker_GetLock_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockLocker_GetLock_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockLocker) List() *MockLocker_List_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "List", params, verifier.timeout)
	return &MockLocker_List_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockLocker_List_OngoingVerification struct {
	mock              *MockLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockLocker_List_OngoingVerification) GetCapturedArguments() {
}

func (c *MockLocker_List_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockLocker) TryLock(_param0 models.Project, _param1 string, _param2 models.PullRequest, _param3 models.User) *MockLocker_TryLock_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2, _param3}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "TryLock", params, verifier.timeout)
	return &MockLocker_TryLock_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockLocker_TryLock_OngoingVerification struct {
	mock              *MockLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockLocker_TryLock_OngoingVerification) GetCapturedArguments() (models.Project, string, models.PullRequest, models.User) {
	_param0, _param1, _param2, _param3 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1], _param3[len(_param3)-1]
}

func (c *MockLocker_TryLock_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Project, _param1 []string, _param2 []models.PullRequest, _param3 []models.User) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Project, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Project)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(models.PullRequest)
		}
		_param3 = make([]models.User, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(models.User)
		}
	}
	return
}

func (verifier *VerifierMockLocker) Unlock(_param0 string) *MockLocker_Unlock_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Unlock", params, verifier.timeout)
	return &MockLocker_Unlock_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockLocker_Unlock_OngoingVerification struct {
	mock              *MockLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockLocker_Unlock_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockLocker_Unlock_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockLocker) UnlockByPull(_param0 string, _param1 int) *MockLocker_UnlockByPull_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UnlockByPull", params, verifier.timeout)
	return &MockLocker_UnlockByPull_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockLocker_UnlockByPull_OngoingVerification struct {
	mock              *MockLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockLocker_UnlockByPull_OngoingVerification) GetCapturedArguments() (string, int) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockLocker_UnlockByPull_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []int) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]int, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(int)
		}
	}
	return
}
