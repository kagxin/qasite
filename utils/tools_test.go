package utils

import "testing"

func TestMd5(t *testing.T) {
	if Md5("123456") != "e10adc3949ba59abbe56e057f20f883e" {
		t.Error("Md5结果错误")
	}
}

func TestGenToken(t *testing.T) {
	token := GenToken(1)
	t.Logf("token: %v", token)
}
