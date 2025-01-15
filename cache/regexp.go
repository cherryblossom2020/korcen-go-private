package cache

import "regexp"

var URL = regexp.MustCompile(`https?://[^\s]+|www\.[^\s]+`)
