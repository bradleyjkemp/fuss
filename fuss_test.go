package fuss

import (
	"crypto/rand"
	"encoding/binary"
	"net/http"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestString(t *testing.T) {
	testString := "Hello world"
	var result string
	Seed(append([]byte{byte(len(testString))}, []byte(testString)...)).Fuss(&result)
	if result != testString {
		t.Fatal("Result", result, "not equal expected", testString)
	}
}

type testStruct struct {
	s string
	i int
}

func TestStruct(t *testing.T) {
	result := testStruct{}
	testString := "Hello world"
	seed := append([]byte{byte(len(testString))}, []byte(testString)...)
	ten := make([]byte, 8)
	binary.BigEndian.PutUint64(ten, 10)
	seed = append(seed, ten...)
	Seed(seed).Fuss(&result)
	expected := testStruct{testString, 10}
	if !reflect.DeepEqual(result, expected) {
		t.Fatal("Result", result, "not equal expected", expected)
	}
}

func TestPointer(t *testing.T) {
	var result *string
	Seed([]byte{1, 3, 'a', 'b', 'c'}).Fuss(&result)
	if result == nil || *result != "abc" {
		t.Fatal("Result", result, "not equal expected")
	}

	result = nil
	Seed([]byte{0, 3, 'a', 'b', 'c'}).Fuss(&result)
	if result != nil {
		t.Fatal("Result", result, "not equal expected (nil)")
	}
}

func TestHttpRequest(t *testing.T) {
	data := make([]byte, 1024)
	rand.Read(data)
	r := http.Request{}
	Seed(data).Fuss(&r)
	spew.Dump(r)
}
