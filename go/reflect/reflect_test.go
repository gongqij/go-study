package main

import "testing"

func TestAll(t *testing.T) {
	t.Run("TestModifySliceByReflect", TestModifySliceByReflect)
	t.Run("TestDeepEqual", TestDeepEqual)
	t.Run("TestImplement", TestImplement)
	t.Run("TestAddByReflect", TestAddByReflect)
	t.Run("TestPrintMethodAndField", TestPrintMethodAndField)
}

func TestModifySliceByReflect(t *testing.T) {
	modifySliceByReflect()
}

func TestDeepEqual(t *testing.T) {
	deepEqual()
}

func TestImplement(t *testing.T) {
	implement()
}

func TestAddByReflect(t *testing.T) {
	addByReflect()
}

func TestPrintMethodAndField(t *testing.T) {
	printMethodAndField()
}
