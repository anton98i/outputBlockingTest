package main

import (
	"testing"
	"time"
)

func TestTestFunctions(t *testing.T) {
	beforeTimerCalled := 0
	afterTimerCalled := 0

	beforeCalled := 0
	testCalled := 0
	currentIterationAfterCalls := 0
	lastIterationAfterCAlls := 0
	afterCalled := 0

	var millisecondsToTest int64 = 1

	beforeCallAtAfter := 0
	maxDuration := time.Duration(0)
	duration := time.Duration(0)
	var startTime time.Time
	var endTime time.Time

	testFunctions([]testFunction{
		testFunction{
			name: "test",
			init: func() testFunctionInstance {
				return testFunctionInstance{
					beforeTimer: func() {
						beforeTimerCalled++
						beforeTimerCalled++
						startTime = time.Now()
						if afterTimerCalled > 0 {
							pauseTime := startTime.Sub(endTime)
							if pauseTime < duration {
								t.Errorf("Pause during testFunction is too small, only %d ms", pauseTime)
							}
						}
					},
					before: func() {
						beforeCalled++
						currentIterationAfterCalls = 0
					},
					testFunc: func() {
						if beforeCallAtAfter > beforeTimerCalled {
							beforeCallAtAfter = beforeTimerCalled
							t.Errorf("testFunc called, but not before")
						}
						testCalled++
						currentIterationAfterCalls++
					},
					after: func() {
						beforeCallAtAfter = beforeTimerCalled
						afterCalled++
						if beforeTimerCalled < afterCalled {
							t.Errorf("beforeTimerCalled < afterCalled (%d < %d)", beforeTimerCalled, afterCalled)
						}
						if lastIterationAfterCAlls > currentIterationAfterCalls {
							t.Errorf("lastIterationAfterCAlls > currentIterationAfterCalls (%d > %d)", lastIterationAfterCAlls, currentIterationAfterCalls)
						}
						lastIterationAfterCAlls = currentIterationAfterCalls
					},
					afterTimer: func() {
						afterTimerCalled++
						if beforeTimerCalled < afterTimerCalled {
							t.Errorf("beforeTimerCalled < afterTimerCalled (%d < %d)", beforeTimerCalled, afterTimerCalled)
						}
						endTime = time.Now()
						duration = endTime.Sub(startTime)
						if duration > maxDuration {
							maxDuration = duration
						}
					},
				}
			},
		},
	}, millisecondsToTest, false)

	if beforeCalled == 0 {
		t.Errorf("before not called")
	}
	if testCalled == 0 {
		t.Errorf("testFunc not called")
	}
	if afterCalled == 0 {
		t.Errorf("after not called")
	}
	if beforeTimerCalled == 0 {
		t.Errorf("beforeTimerCalled not called")
	}
	if afterTimerCalled == 0 {
		t.Errorf("afterTimerCalled not called")
	}
	if maxDuration.Milliseconds() < millisecondsToTest {
		t.Errorf("Milliseconds not big enough: %d < %d", maxDuration.Milliseconds(), millisecondsToTest)
	}
}
