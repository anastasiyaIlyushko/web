package status

//import "gopkg.in/webnice/debug.v1"
//import "gopkg.in/webnice/log.v2"
import (
	"testing"
)

func TestStatus(tst *testing.T) {
	var v int
	var t string
	var b []byte

	v = Ok
	t = Text(v)
	b = Bytes(v)

	if v != 200 {
		tst.Errorf("Error constant")
	}
	if t != "Ok" {
		tst.Errorf("Error constant")
	}
	if string(b) != t {
		tst.Errorf("Error constant")
	}

	t = Text(999)
	b = Bytes(999)
	if t != string(b) || t != `HTTP status code 999` {
		tst.Errorf("Error constant")
	}
}
