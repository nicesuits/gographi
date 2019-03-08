package gographi

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// Video struct
type Video struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	User        User      `json:"user"`
	URL         string    `json:"url"`
	CreatedAt   time.Time `json:"createdAt"`
	Related     []Video
}

// MarshalID returns a marshaler for ints
func MarshalID(id int) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(fmt.Sprintf("%d", id)))
	})
}

// UnmarshalID returns an integer from a strings
func UnmarshalID(v interface{}) (int, error) {
	if id, ok := v.(string); ok {
		i, e := strconv.Atoi(id)
		return int(i), e
	}
	return 0, fmt.Errorf("ids must be strings")
}

// MarshalTimestamp returns
func MarshalTimestamp(t time.Time) graphql.Marshaler {
	timestamp := t.Unix() * 1000
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(timestamp, 10))
	})
}

// UnmarshalTimestamp retunrs
func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(int); ok {
		return time.Unix(int64(tmpStr), 0), nil
	}
	return time.Time{}, errors.TimeStampError
}
