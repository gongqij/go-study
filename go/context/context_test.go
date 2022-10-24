package main

import "testing"

func TestListeningForCancellation(t *testing.T) {
	listeningForCancellation()
}

func TestEmittingCancellationEvent(t *testing.T) {
	emittingCancellationEvent()
}

func TestContextTimeout(t *testing.T) {
	contextTimeout()
}
