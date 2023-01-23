// Code generated by "fsmfy.go"; DO NOT EDIT.
package main

import (
	"context"

	"github.com/looplab/fsm"
)

type (
	// TestTypeState represents all states for this FSM.
	TestTypeState string
	//TestTypeEvent  represents all events for this FSM.
	TestTypeEvent string
)

const (
	// States.
	// IdleState is initial state.
	IdleState TestTypeState = "idle"

	// ArchivedState state.
	ArchivedState TestTypeState = "archived"
	// AwaitingPaymentState state.
	AwaitingPaymentState TestTypeState = "awaiting_payment"
	// ConfirmedState state.
	ConfirmedState TestTypeState = "confirmed"
	// DisputeState state.
	DisputeState TestTypeState = "dispute"
	// CompletedState state.
	CompletedState TestTypeState = "completed"
	// CreatedState state.
	CreatedState TestTypeState = "created"
	// CanceledState state.
	CanceledState TestTypeState = "canceled"

	//Events.

	// ResolveEvent state.
	ResolveEvent TestTypeEvent = "RESOLVE"
	// SendMessageEvent state.
	SendMessageEvent TestTypeEvent = "SEND_MESSAGE"
	// DisputeEvent state.
	DisputeEvent TestTypeEvent = "DISPUTE"
	// AutoCompleteOnTimeoutEvent state.
	AutoCompleteOnTimeoutEvent TestTypeEvent = "AUTO_COMPLETE_ON_TIMEOUT"
	// MusoCancelEvent state.
	MusoCancelEvent TestTypeEvent = "MUSO_CANCEL"
	// CancelEvent state.
	CancelEvent TestTypeEvent = "CANCEL"
	// CompleteEvent state.
	CompleteEvent TestTypeEvent = "COMPLETE"
	// MusoConfirmEvent state.
	MusoConfirmEvent TestTypeEvent = "MUSO_CONFIRM"
	// ArchiveEvent state.
	ArchiveEvent TestTypeEvent = "ARCHIVE"
	// BookAGigEvent state.
	BookAGigEvent TestTypeEvent = "BOOK_A_GIG"
	// SetPriceEvent state.
	SetPriceEvent TestTypeEvent = "SET_PRICE"
	// PaySuccessfulEvent state.
	PaySuccessfulEvent TestTypeEvent = "PAY_SUCCESSFUL"
	// PaymentFailedEvent state.
	PaymentFailedEvent TestTypeEvent = "PAYMENT_FAILED"
	// CancelBeforeThresholdEvent state.
	CancelBeforeThresholdEvent TestTypeEvent = "CANCEL_BEFORE_THRESHOLD"
)

// TestType  implements FSM.
type TestType struct {
	fsm *fsm.FSM
}

// NewTestType creates new FSM with callbacks provided.
func NewTestType(callbacks fsm.Callbacks) *TestType {
	fsm := fsm.NewFSM(
		IdleState.String(),
		fsm.Events{

			{Name: ResolveEvent.String(), Src: []string{DisputeState.String()}, Dst: CompletedState.String()},
			{Name: SendMessageEvent.String(), Src: []string{CreatedState.String()}, Dst: CreatedState.String()},
			{Name: DisputeEvent.String(), Src: []string{ConfirmedState.String()}, Dst: DisputeState.String()},
			{Name: AutoCompleteOnTimeoutEvent.String(), Src: []string{ConfirmedState.String()}, Dst: CompletedState.String()},
			{Name: MusoCancelEvent.String(), Src: []string{ConfirmedState.String()}, Dst: CanceledState.String()},
			{Name: CancelEvent.String(), Src: []string{AwaitingPaymentState.String(), CreatedState.String()}, Dst: CanceledState.String()},
			{Name: CompleteEvent.String(), Src: []string{ConfirmedState.String()}, Dst: CompletedState.String()},
			{Name: MusoConfirmEvent.String(), Src: []string{CreatedState.String()}, Dst: AwaitingPaymentState.String()},
			{Name: ArchiveEvent.String(), Src: []string{CompletedState.String(), CanceledState.String()}, Dst: ArchivedState.String()},
			{Name: BookAGigEvent.String(), Src: []string{IdleState.String()}, Dst: CreatedState.String()},
			{Name: SetPriceEvent.String(), Src: []string{CreatedState.String()}, Dst: CreatedState.String()},
			{Name: PaySuccessfulEvent.String(), Src: []string{AwaitingPaymentState.String()}, Dst: ConfirmedState.String()},
			{Name: PaymentFailedEvent.String(), Src: []string{AwaitingPaymentState.String()}, Dst: AwaitingPaymentState.String()},
			{Name: CancelBeforeThresholdEvent.String(), Src: []string{ConfirmedState.String()}, Dst: CanceledState.String()},
		},
		callbacks,
	)
	return &TestType{fsm: fsm}
}

// String returns string representation of the state.
func (s TestTypeState) String() string {
	return string(s)
}

// String returns string representation of the event.
func (s TestTypeEvent) String() string {
	return string(s)
}

// Current returns the current state of the TestType.
func (f *TestType) Current() TestTypeState {
	return TestTypeState(f.fsm.Current())
}

// Is returns true if state is the current state.
func (f *TestType) Is(state TestTypeState) bool {
	return f.fsm.Is(state.String())
}

// SetState allows the user to move to the given state from current state.
// The call does not trigger any callbacks, if defined.
func (f *TestType) SetState(state TestTypeState) {
	f.fsm.SetState(state.String())
}

// Can returns true if event can occur in the current state.
func (f *TestType) Can(event TestTypeEvent) bool {
	return f.fsm.Can(event.String())
}

// AvailableTransitions returns a list of transitions available in the
// current state.
func (f *TestType) AvailableTransitions() []string {
	return f.fsm.AvailableTransitions()
}

// Cannot returns true if event can not occur in the current state.
// It is a convenience method to help code read nicely.
func (f *TestType) Cannot(event TestTypeEvent) bool {
	return f.fsm.Cannot(event.String())
}

// Metadata returns the value stored in metadata
func (f *TestType) Metadata(key string) (interface{}, bool) {
	return f.fsm.Metadata(key)
}

// SetMetadata stores the dataValue in metadata indexing it with key
func (f *TestType) SetMetadata(key string, dataValue interface{}) {
	f.fsm.SetMetadata(key, dataValue)
}

// DeleteMetadata deletes the dataValue in metadata by key
func (f *TestType) DeleteMetadata(key string) {
	f.fsm.DeleteMetadata(key)
}

// Event initiates a state transition with the named event.
//
// The call takes a variable number of arguments that will be passed to the
// callback, if defined.
func (f *TestType) Event(ctx context.Context, event TestTypeEvent, args ...interface{}) error {
	return f.fsm.Event(ctx, event.String(), args...)
}