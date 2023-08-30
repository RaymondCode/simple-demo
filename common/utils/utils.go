package utils

import (
	"google.golang.org/grpc/status"
	"regexp"
)

func MatchError(err error, target *status.Status) (*status.Status, bool) {
	st, _ := status.FromError(err)
	if st.Message() == target.Message() {
		return st, true
	}
	return st, false
}

func MatchRegexp(pattern string, value string) bool {
	r := regexp.MustCompile(pattern)
	return r.MatchString(value)
}
