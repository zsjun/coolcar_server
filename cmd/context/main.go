package main

import (
	"context"
	"fmt"
	"time"
)

type parameKey struct {

}
func main() {
	c := context.WithValue(context.Background(), parameKey{}, "abc");
	c, cancel := context.WithTimeout(c, 5 * time.Second);
	defer cancel();
	mainTask(c)
}

func mainTask(c context.Context) {
	fmt.Printf("startred with params %q\n", c.Value(parameKey{}))
	smallTask(c, "task1",4 * time.Second);
	smallTask(c, "task2", 2  *time.Second);
}

func smallTask(c context.Context, name string, d time.Duration) {
	fmt.Printf("%s startred %q\n", name,c.Value(parameKey{}))
	select {
		case <-time.After(d):
			fmt.Printf("%s done \n", name);
		case <-c.Done():
			fmt.Printf("%s canceled \n", name);
	}
	// <-c.Done();
}