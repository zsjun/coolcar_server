package main

import (
	trippb "coolcar/proto/gen/go"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func main() {
	trip := trippb.Trip{
		Start: "abc",
		End: "sd",
		DurationSec: 3600,
		FeeCent: 10000,
	}
	fmt.Println(&trip)
  b,err :=	proto.Marshal(&trip)
	if(err != nil) {
		panic(err)
	}
	fmt.Printf("%X\n",b);

}