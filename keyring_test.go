package keyctl

import (
	"testing"
)

func TestAdd100BytesToSessionKeyring(t *testing.T) {
	ring, err := SessionKeyring()
	if err != nil {
		t.Fatal(err)
	}
	key, err := ring.Add("blank", make([]byte, 100))
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("added 100 byte empty key as: %v\n", key.Id())

	buf, err := key.Get()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("read %d octets from key: %v\n", len(buf), buf)
}

func TestFetchKeyFail(t *testing.T) {
	ring, err := SessionKeyring()
	if err != nil {
		t.Fatal(err)
	}
	_, err = ring.Search("abigbunchofnonsense")
	if err == nil {
		t.Fatal("search expected to fail")
	}
}

func TestUnlinkKey(t *testing.T) {
	ring, err := SessionKeyring()
	if err != nil {
		t.Fatal(err)
	}
	key, err := ring.Search("blank")
	if err != nil {
		t.Fatal(err)
	}
	if err = key.Unlink(); err != nil {
		t.Fatal(err)
	}
}
