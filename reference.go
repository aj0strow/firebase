package firebase

import (
	"strings"
)

// Reference is a Firebase URL path.
type Reference []string

// Root returns a reference to "/".
func (ref Reference) Root() Reference {
	return nil
}

// Parent returns the URL path one segment up. Eg: "/users/aj0strow" parent is "/users".
func (ref Reference) Parent() Reference {
	if len(ref) == 0 {
		return nil
	}
	return ref[0 : len(ref)-1]
}

// Child returns the URL path one or more segments down. Eg: "/users" child could be "/users/aj0strow/profile".
func (ref Reference) Child(more ...string) Reference {
	return append(ref, more...)
}

// Push returns a child reference with a new Firebase Push ID.
func (ref Reference) Push() Reference {
	return ref.Child(NewPushID())
}

// String returns the URL path as a slash-separated string.
func (ref Reference) String() string {
	return "/" + strings.Join(ref, "/")
}
