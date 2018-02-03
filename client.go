package firebase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

var ErrEmptyPath = errors.New("firebase: reference has empty path segment")
var ErrInvalidDatabaseURL = errors.New("firebase: invalid database URL")

// Client interfaces with one Database URL and one auth context.
type Client struct {
	client      *http.Client
	databaseURL string
	auth        string
}

// Params are type safe Firebase Query Parameters in the REST API.
type Params struct {
	OrderBy      string
	EqualTo      string
	StartAt      string
	EndAt        string
	LimitToFirst int
	LimitToLast  int
	Shallow      bool
}

// Query fetches and parses JSON data at the reference path.
func (c *Client) Query(ref Reference, params *Params, v interface{}) error {
	url, err := c.loc(ref, params)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	w, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer w.Body.Close()
	return json.NewDecoder(w.Body).Decode(v)
}

// Write encodes and overwrites JSON data at the reference path.
func (c *Client) Write(ref Reference, v interface{}) error {
	body, err := json.Marshal(v)
	if err != nil {
		return err
	}
	url, err := c.loc(ref, nil)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	w, err := c.client.Do(req)
	if err != nil {
		return err
	}
	if w.StatusCode != 200 {
		return fmt.Errorf("firebase: write failed with %d", w.Status)
	}
	return nil
}

// UpdateByMerge encodes and deep merges JSON data at the root path.
func (c *Client) UpdateByMerge(batch *Batch) error {
	update, err := batch.Merge()
	if err != nil {
		return err
	}
	body, err := json.Marshal(update)
	if err != nil {
		return err
	}
	url, err := c.loc(nil, nil)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PATCH", url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	w, err := c.client.Do(req)
	if err != nil {
		return err
	}
	if w.StatusCode != 200 {
		return fmt.Errorf("firebase: write failed with %d", w.Status)
	}
	return nil
}

// Remove deletes data at the reference path.
func (c *Client) Remove(ref Reference) error {
	url, err := c.loc(ref, nil)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	w, err := c.client.Do(req)
	if err != nil {
		return err
	}
	if w.StatusCode != 200 {
		return fmt.Errorf("firebase: remove failed with %d", w.Status)
	}
	return nil
}

func (c *Client) loc(ref Reference, params *Params) (string, error) {
	for _, segment := range ref {
		if segment == "" {
			return "", ErrEmptyPath
		}
	}
	rem, err := url.Parse(c.databaseURL)
	if err != nil {
		return "", ErrInvalidDatabaseURL
	}
	rel, err := url.Parse(ref.String() + ".json")
	if err != nil {
		return "", err
	}
	loc := rem.ResolveReference(rel)
	qs := url.Values{}
	if c.auth != "" {
		qs.Set("auth", c.auth)
	}
	if params != nil {
		if params.OrderBy != "" {
			qs.Set("orderBy", params.OrderBy)
		}
		if params.EqualTo != "" {
			qs.Set("equalTo", params.EqualTo)
		}
		if params.StartAt != "" {
			qs.Set("startAt", params.StartAt)
		}
		if params.EndAt != "" {
			qs.Set("endAt", params.EndAt)
		}
		if params.LimitToFirst > 0 {
			qs.Set("limitToFirst", strconv.Itoa(params.LimitToFirst))
		}
		if params.LimitToLast > 0 {
			qs.Set("limitToLast", strconv.Itoa(params.LimitToLast))
		}
		if params.Shallow {
			qs.Set("shallow", "true")
		}
	}
	loc.RawQuery = qs.Encode()
	return loc.String(), nil
}
