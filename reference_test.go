package firebase

import (
	"testing"
)

func TestReference(t *testing.T) {
	root := Reference{}
	if root.Join() != "/" {
		t.Errorf("root reference broken")
	}
	users := root.Child("users")
	if users.Join() != "/users" {
		t.Errorf("child path broken")
	}
	if users.Parent().Join() != root.Join() {
		t.Errorf("parent path is unexpected")
	}
	newbie := users.Push()
	if len(newbie) != 2 {
		t.Errorf("push ID does not work")
	}
	myFace := users.Root().Child("images", "aj0strow", "profileHttps")
	if len(myFace) != 3 {
		t.Errorf("unexpected traversal")
	}
}
