// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/grishagavrin/auth-chat-grpc/internal/service.NoteService -o note_service_minimock.go -n NoteServiceMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/grishagavrin/auth-chat-grpc/internal/model"
)

// NoteServiceMock implements service.NoteService
type NoteServiceMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreate          func(ctx context.Context, info *model.NoteInfo) (i1 int64, err error)
	inspectFuncCreate   func(ctx context.Context, info *model.NoteInfo)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mNoteServiceMockCreate

	funcGet          func(ctx context.Context, id int64) (np1 *model.Note, err error)
	inspectFuncGet   func(ctx context.Context, id int64)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mNoteServiceMockGet
}

// NewNoteServiceMock returns a mock for service.NoteService
func NewNoteServiceMock(t minimock.Tester) *NoteServiceMock {
	m := &NoteServiceMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mNoteServiceMockCreate{mock: m}
	m.CreateMock.callArgs = []*NoteServiceMockCreateParams{}

	m.GetMock = mNoteServiceMockGet{mock: m}
	m.GetMock.callArgs = []*NoteServiceMockGetParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mNoteServiceMockCreate struct {
	mock               *NoteServiceMock
	defaultExpectation *NoteServiceMockCreateExpectation
	expectations       []*NoteServiceMockCreateExpectation

	callArgs []*NoteServiceMockCreateParams
	mutex    sync.RWMutex
}

// NoteServiceMockCreateExpectation specifies expectation struct of the NoteService.Create
type NoteServiceMockCreateExpectation struct {
	mock    *NoteServiceMock
	params  *NoteServiceMockCreateParams
	results *NoteServiceMockCreateResults
	Counter uint64
}

// NoteServiceMockCreateParams contains parameters of the NoteService.Create
type NoteServiceMockCreateParams struct {
	ctx  context.Context
	info *model.NoteInfo
}

// NoteServiceMockCreateResults contains results of the NoteService.Create
type NoteServiceMockCreateResults struct {
	i1  int64
	err error
}

// Expect sets up expected params for NoteService.Create
func (mmCreate *mNoteServiceMockCreate) Expect(ctx context.Context, info *model.NoteInfo) *mNoteServiceMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("NoteServiceMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &NoteServiceMockCreateExpectation{}
	}

	mmCreate.defaultExpectation.params = &NoteServiceMockCreateParams{ctx, info}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the NoteService.Create
func (mmCreate *mNoteServiceMockCreate) Inspect(f func(ctx context.Context, info *model.NoteInfo)) *mNoteServiceMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for NoteServiceMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by NoteService.Create
func (mmCreate *mNoteServiceMockCreate) Return(i1 int64, err error) *NoteServiceMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("NoteServiceMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &NoteServiceMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &NoteServiceMockCreateResults{i1, err}
	return mmCreate.mock
}

// Set uses given function f to mock the NoteService.Create method
func (mmCreate *mNoteServiceMockCreate) Set(f func(ctx context.Context, info *model.NoteInfo) (i1 int64, err error)) *NoteServiceMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the NoteService.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the NoteService.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the NoteService.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mNoteServiceMockCreate) When(ctx context.Context, info *model.NoteInfo) *NoteServiceMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("NoteServiceMock.Create mock is already set by Set")
	}

	expectation := &NoteServiceMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &NoteServiceMockCreateParams{ctx, info},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up NoteService.Create return parameters for the expectation previously defined by the When method
func (e *NoteServiceMockCreateExpectation) Then(i1 int64, err error) *NoteServiceMock {
	e.results = &NoteServiceMockCreateResults{i1, err}
	return e.mock
}

// Create implements service.NoteService
func (mmCreate *NoteServiceMock) Create(ctx context.Context, info *model.NoteInfo) (i1 int64, err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, info)
	}

	mm_params := NoteServiceMockCreateParams{ctx, info}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, &mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_got := NoteServiceMockCreateParams{ctx, info}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("NoteServiceMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the NoteServiceMock.Create")
		}
		return (*mm_results).i1, (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, info)
	}
	mmCreate.t.Fatalf("Unexpected call to NoteServiceMock.Create. %v %v", ctx, info)
	return
}

// CreateAfterCounter returns a count of finished NoteServiceMock.Create invocations
func (mmCreate *NoteServiceMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of NoteServiceMock.Create invocations
func (mmCreate *NoteServiceMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to NoteServiceMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mNoteServiceMockCreate) Calls() []*NoteServiceMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*NoteServiceMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *NoteServiceMock) MinimockCreateDone() bool {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateInspect logs each unmet expectation
func (m *NoteServiceMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to NoteServiceMock.Create with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to NoteServiceMock.Create")
		} else {
			m.t.Errorf("Expected call to NoteServiceMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		m.t.Error("Expected call to NoteServiceMock.Create")
	}
}

type mNoteServiceMockGet struct {
	mock               *NoteServiceMock
	defaultExpectation *NoteServiceMockGetExpectation
	expectations       []*NoteServiceMockGetExpectation

	callArgs []*NoteServiceMockGetParams
	mutex    sync.RWMutex
}

// NoteServiceMockGetExpectation specifies expectation struct of the NoteService.Get
type NoteServiceMockGetExpectation struct {
	mock    *NoteServiceMock
	params  *NoteServiceMockGetParams
	results *NoteServiceMockGetResults
	Counter uint64
}

// NoteServiceMockGetParams contains parameters of the NoteService.Get
type NoteServiceMockGetParams struct {
	ctx context.Context
	id  int64
}

// NoteServiceMockGetResults contains results of the NoteService.Get
type NoteServiceMockGetResults struct {
	np1 *model.Note
	err error
}

// Expect sets up expected params for NoteService.Get
func (mmGet *mNoteServiceMockGet) Expect(ctx context.Context, id int64) *mNoteServiceMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("NoteServiceMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &NoteServiceMockGetExpectation{}
	}

	mmGet.defaultExpectation.params = &NoteServiceMockGetParams{ctx, id}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the NoteService.Get
func (mmGet *mNoteServiceMockGet) Inspect(f func(ctx context.Context, id int64)) *mNoteServiceMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for NoteServiceMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by NoteService.Get
func (mmGet *mNoteServiceMockGet) Return(np1 *model.Note, err error) *NoteServiceMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("NoteServiceMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &NoteServiceMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &NoteServiceMockGetResults{np1, err}
	return mmGet.mock
}

// Set uses given function f to mock the NoteService.Get method
func (mmGet *mNoteServiceMockGet) Set(f func(ctx context.Context, id int64) (np1 *model.Note, err error)) *NoteServiceMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the NoteService.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the NoteService.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the NoteService.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mNoteServiceMockGet) When(ctx context.Context, id int64) *NoteServiceMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("NoteServiceMock.Get mock is already set by Set")
	}

	expectation := &NoteServiceMockGetExpectation{
		mock:   mmGet.mock,
		params: &NoteServiceMockGetParams{ctx, id},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up NoteService.Get return parameters for the expectation previously defined by the When method
func (e *NoteServiceMockGetExpectation) Then(np1 *model.Note, err error) *NoteServiceMock {
	e.results = &NoteServiceMockGetResults{np1, err}
	return e.mock
}

// Get implements service.NoteService
func (mmGet *NoteServiceMock) Get(ctx context.Context, id int64) (np1 *model.Note, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(ctx, id)
	}

	mm_params := NoteServiceMockGetParams{ctx, id}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, &mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.np1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_got := NoteServiceMockGetParams{ctx, id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("NoteServiceMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the NoteServiceMock.Get")
		}
		return (*mm_results).np1, (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(ctx, id)
	}
	mmGet.t.Fatalf("Unexpected call to NoteServiceMock.Get. %v %v", ctx, id)
	return
}

// GetAfterCounter returns a count of finished NoteServiceMock.Get invocations
func (mmGet *NoteServiceMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of NoteServiceMock.Get invocations
func (mmGet *NoteServiceMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to NoteServiceMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mNoteServiceMockGet) Calls() []*NoteServiceMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*NoteServiceMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *NoteServiceMock) MinimockGetDone() bool {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetInspect logs each unmet expectation
func (m *NoteServiceMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to NoteServiceMock.Get with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to NoteServiceMock.Get")
		} else {
			m.t.Errorf("Expected call to NoteServiceMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to NoteServiceMock.Get")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *NoteServiceMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateInspect()

			m.MinimockGetInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *NoteServiceMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *NoteServiceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockGetDone()
}