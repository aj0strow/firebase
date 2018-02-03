package firebase

import (
	"strings"
)

type Reference []string

func (ref Reference) Root() Reference {
	return nil
}

func (ref Reference) Parent() Reference {
	if len(ref) == 0 {
		return nil
	}
	return ref[0 : len(ref)-1]
}

func (ref Reference) Child(more ...string) Reference {
	return append(ref, more...)
}

func (ref Reference) Push() Reference {
	return ref.Child(NewPushID())
}

func (ref Reference) Join() string {
	return "/" + strings.Join(ref, "/")
}
