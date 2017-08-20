package main

import (
	"fmt"
	"hash/crc32"
	"regexp"
	"testing"
)

func TestRandSecure(t *testing.T) {
	num, err := cryptoRandSecure(int64(^(uint64(1) << 63)))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(num)
}

func TestRandSecure_Limit_Max_1(t *testing.T) {
	num, err := cryptoRandSecure(1)
	if err != nil {
		t.Fatal(err)
	}
	if num > 1 {
		t.Error("didn't expect nimber bigger then 1")
	}
}

func TestUniqueRecipeID_No_Title(t *testing.T) {
	id, err := uniqueRecipeID("")
	if err == nil {
		t.Fatal("error was expected")
		return
	}
	if err != errRecipeTitleNotSet {
		t.Error("expected '%s' but got '%s'", errRecipeTitleNotSet, err)
	}
	if id != "" {
		t.Error("in case of error empty id is expected")
	}
}

func TestUniqueRecipeID(t *testing.T) {
	title := "raspberry pi"
	id, err := uniqueRecipeID(title)
	if err != nil {
		t.Fatal("error was expected")
		return
	}
	t.Log(id)

	// use regex to match id
	var crc uint32 = crc32.ChecksumIEEE([]byte(title))
	r, err := regexp.Compile(fmt.Sprintf("%x-[0-9 a-z].*-[0-9 a-z].*", crc))
	if err != nil {
		t.Fatal(err)
	}
	if !r.MatchString(id) {
		t.Error("'%s' is not in expected format", id)
	}
}
