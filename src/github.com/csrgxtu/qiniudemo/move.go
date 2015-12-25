package main

import (
	"fmt"
	"os"

	"golang.org/x/net/context"
	"qiniupkg.com/api.v7/kodo"
)

func main() {
	argsWithoutProg := os.Args[1:]

	kodo.SetMac("_4TUdWfMQGZ5f2DFFmXbARs7pQLWmiPK-IFbSsw5", "1x0lUvV11qxbWQO1G_XrMm6v-MSsDWJWNCJk2K67")
	zone := 0
	c := kodo.New(zone, nil)
	bucket := c.Bucket("brpublic")
	ctx := context.Background()
	from := argsWithoutProg[0]
	to := argsWithoutProg[1]
	err := bucket.Move(ctx, "rigo/"+from, "rigo/"+to)
	//err := bucket.Delete(ctx, "rigo/"+localFile)
	//err := bucket.PutFile(ctx, nil, "rigo/"+localFile, localFile, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("move success")
		fmt.Println("http://7xj2i2.com2.z0.glb.qiniucdn.com/rigo/" + to)
	}
}
