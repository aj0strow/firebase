package firebase

import (
	"testing"
)

func TestClientLoc(t *testing.T) {
	client := &Client{
		databaseURL: "https://test.firebaseio.com/",
		auth:        "sekret",
	}
	tests := []struct {
		ref    Reference
		params *Params
		url    string
	}{
		{
			ref:    nil,
			params: nil,
			url:    "https://test.firebaseio.com/.json?auth=sekret",
		},
		{
			ref: Reference{"users"},
			params: &Params{
				OrderBy: "name",
				StartAt: "a",
				EndAt:   "a~",
			},
			url: "https://test.firebaseio.com/users.json?auth=sekret&endAt=a~&orderBy=name&startAt=a",
		},
		{
			ref: Reference{"users"},
			params: &Params{
				Shallow:      true,
				LimitToFirst: 1000,
			},
			url: "https://test.firebaseio.com/users.json?auth=sekret&limitToFirst=1000&shallow=true",
		},
		{
			ref: Reference{"users"},
			params: &Params{
				OrderBy: "region",
				EqualTo: "us",
			},
			url: "https://test.firebaseio.com/users.json?auth=sekret&equalTo=us&orderBy=region",
		},
	}
	for _, tt := range tests {
		url, err := client.loc(tt.ref, tt.params)
		if err != nil {
			t.Fatal(err)
		}
		if url != tt.url {
			t.Errorf("have %s, want %s", url, tt.url)
		}
	}
}
