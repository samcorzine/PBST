package main

import (
	"testing"
)

func TestMoveChecker(t *testing.T) {
	testTree := Tree{Root: nil}
	testTree.put("test", 1)
	result, err := testTree.get("test")
	if result != 1 || err != nil {
		t.Errorf("failed")
	}
	testTree.put("test2", 2)
	result, err = testTree.get("test2")
	if result != 2 || err != nil {
		t.Errorf("failed to get second key")
	}
	testTree.put("tesa", 0)
	result, err = testTree.get("tesa")
	if result != 0 || err != nil {
		t.Errorf("failed")
	}
	testTree.put("test", 2)
	result, err = testTree.get("test")
	if result != 2 || err != nil {
		t.Errorf("failed")
	}
	result, err = testTree.get("gsasdgasdg")
	if result != -1 || err.Error() != "Invalid Key: gsasdgasdg" {
		t.Errorf("failed")
	}
	testTree.put("test3", 1423)
	testTree.put("test4", 12)
	testTree.put("test5", 7)
	testTree.put("test6", 9)
	testTree.put("test7", 4)
	testTree.put("test8", 12341243)
	result, err = testTree.get("test8")
	if result != 12341243 || err != nil {
		t.Errorf("failed")
	}
	testTree.delete("test")
	testTree.delete("test7")
	result, err = testTree.get("test7")
	if result != -1 || err == nil {
		t.Errorf("failed")
		t.Errorf(err.Error())
	}
	result, err = testTree.get("test")
	if result != -1 || err == nil {
		t.Errorf("failed")
	}
}
