package util

import "time"

func RandomDuration() time.Duration {
	return time.Duration(RandomInt(1, 300)) * time.Second
}

func RandomKeyMap() map[string]string {
	m := make(map[string]string)
	n := int(RandomInt(1, 5))
	for i := 0; i < n; i++ {
		m[RandomString(10)] = RandomString(10)
	}
	return m
}
