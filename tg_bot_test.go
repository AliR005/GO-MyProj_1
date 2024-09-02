package main

import "testing"

func TestGetMd5(t *testing.T) {
	hash_is_abc := "900150983cd24fb0d6963f7d28e17f72"
	hash_is_my_func := GetMd5("abc")

	if hash_is_abc != hash_is_my_func {
		t.Errorf("Result was incorrect, got: %s, want: %s", hash_is_abc, hash_is_my_func)
	}
}
func TestAppendMap(t *testing.T) {
	hash_first_el := "0cc175b9c0f1b6a831c399e269772661"
	value_first_el := "hello world"
	append_map()

	if map_hashs[hash_first_el] != value_first_el {
		t.Errorf("Result was incorrect")
	}
}

func TestSettingBot_1(t *testing.T) {
	pass := "c"
	val := "imac"
	if settings_bot(pass) != val {
		t.Errorf("Result was incorrect")
	}
}

func TestSettingBot_2(t *testing.T) {
	key := "0cc175b9c0f1b6a831c399e269772661"
	val := "hello world"
	if settings_bot(val) != key {
		t.Errorf("Result was incorrect")
	}
}
