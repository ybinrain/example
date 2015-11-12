package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"reflect"
)

func main() {
	c, err := redis.Dial("tcp", ":6380")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	fmt.Println(reflect.TypeOf(c))

	r, err := c.Do("PING")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)

	fmt.Println(c.Do("SET", "test", 10))

	c.Do("foo", 1)
	//exists, _ := redis.Int64(c.Do("exists", "foo"))
	//fmt.Println(reflect.TypeOf(exists))
	//fmt.Println(exists)
	//fmt.Printf("%#v\n", exists)
	c.Send("HMSET", "album:1", "title", "Red", "rating", 5)
	c.Send("HMSET", "album:2", "title", "Earthbound", "rating", 1)
	c.Send("HMSET", "album:3", "title", "Beat")
	c.Send("LPUSH", "albums", "1")
	c.Send("LPUSH", "albums", "2")
	c.Send("LPUSH", "albums", "3")
	values, err := redis.Values(c.Do("SORT", "albums",
		"BY", "album:*->rating",
		"GET", "album:*->title",
		"GET", "album:*->rating"))
	if err != nil {
		panic(err)
	}

	for len(values) > 0 {
		var title string
		rating := -1 // initialize to illegal value to detect nil.
		values, err = redis.Scan(values, &title, &rating)
		if err != nil {
			panic(err)
		}
		if rating == -1 {
			fmt.Println(title, "not-rated")
		} else {
			fmt.Println(title, rating)
		}
	}
}
