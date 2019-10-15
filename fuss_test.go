package fuss

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	testString := "Hello world"
	var result string
	Seed(append([]byte{byte(len(testString))}, []byte(testString)...)).Fuzz(&result)
	if result != testString {
		t.Fatal("Result", result, "not equal expected", testString)
	}
}

type testStruct struct {
	S string
	I int
}

func TestStruct(t *testing.T) {
	result := testStruct{}
	testString := "Hello world"
	seed := append([]byte{byte(len(testString))}, []byte(testString)...)
	ten := make([]byte, 8)
	binary.BigEndian.PutUint64(ten, 10)
	seed = append(seed, ten...)
	Seed(seed).Fuzz(&result)
	expected := testStruct{testString, 10}
	if !reflect.DeepEqual(result, expected) {
		t.Fatal("Result", result, "not equal expected", expected)
	}
}

func TestPointer(t *testing.T) {
	var result *string
	Seed([]byte{1, 3, 'a', 'b', 'c'}).Fuzz(&result)
	if result == nil || *result != "abc" {
		t.Fatal("Result", result, "not equal expected")
	}

	result = nil
	Seed([]byte{0, 3, 'a', 'b', 'c'}).Fuzz(&result)
	if result != nil {
		t.Fatal("Result", result, "not equal expected (nil)")
	}
}

func TestIOReader(t *testing.T) {
	var result io.Reader
	Seed([]byte{5, 0, 0, 0, 0}).Fuzz(&result)
	_, ok := result.(*bytes.Reader)
	if !ok {
		t.Fatal("Expected interface to have been implemented with bytes.Reader")
	}
}

func TestHttpRequest(t *testing.T) {
	data := make([]byte, 1024)
	rand.Read(data)
	r := http.Request{}
	Seed(data).Fuzz(&r)
}
