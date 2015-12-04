package main

import (
	"fmt"

	"golang.org/x/net/context"
	"qiniupkg.com/api.v7/kodo"
)

func main() {
	kodo.SetMac("_4TUdWfMQGZ5f2DFFmXbARs7pQLWmiPK-IFbSsw5", "1x0lUvV11qxbWQO1G_XrMm6v-MSsDWJWNCJk2K67")
	zone := 0
	c := kodo.New(zone, nil)
	bucket := c.Bucket("brpublic")
	ctx := context.Background()
	localFile := "dpp.mp4"
	err := bucket.PutFile(ctx, nil, "rigo/dpp.mp4", localFile, nil)
	if err != nil {
		fmt.Println(err)
	}
}
