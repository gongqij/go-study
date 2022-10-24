package powerOnce

import (
	"fmt"
	"testing"
	"time"
)

func TestOnce(t *testing.T) {
	t.Run("call once.Do success", TestSuccess)
	//t.Run("call once.Do failed", TestFailed)

}

func TestSuccess(t *testing.T) {
	var once Once
	t.Logf("第%d次调用once.Do\n", 1)
	err := once.Do(func() error {
		return nil
	})
	for i := 2; i <= 10 && err != nil; i++ {
		time.Sleep(time.Second)
		err = once.Do(func() error {
			return nil
		})
		t.Logf("第%d次调用once.Do\n", i)
	}
	if once.Done() {
		t.Log("init success!")
	} else {
		t.Log("init failed!")
	}
}

func TestFailed(t *testing.T) {
	var once Once
	t.Logf("第%d次调用once.Do\n", 1)
	err := once.Do(func() error {
		return fmt.Errorf("init error")
	})
	for i := 2; i <= 10 && err != nil; i++ {
		time.Sleep(time.Second)
		err = once.Do(func() error {
			return fmt.Errorf("init error")
		})
		t.Logf("第%d次调用once.Do\n", i)
	}
	if once.Done() {
		t.Log("init success!")
	} else {
		t.Log("init failed!")
	}
}
