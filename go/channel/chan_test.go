package main

import "testing"

func TestChan(t *testing.T) {
	t.Run("sendToNilChan", TestSendToNilChan)
}

func TestSendToNilChan(t *testing.T) {
	sendToNilChan()
}

func TestRecvFromClosedChan(t *testing.T) {
	recvFromClosedChan()
}
