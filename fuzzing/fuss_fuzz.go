package fuzzing

import (
	"net/http"

	"github.com/bradleyjkemp/fuss"
)

func FuzzHttpRequest(data []byte) int {
	r := http.Request{}
	fuss.Seed(data).Fuss(&r)
	return 0
}
