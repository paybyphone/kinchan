# AWS Kinesis Channels

Makes Kinesis streams look like Go channels.

## Usage:

```
package main

import (
	"flag"
	"fmt"
	"log"

	"bitbucket.org/rreppel/kinchan"
)

func main() {
	streamPtr := flag.String("s", "", "Kinesis stream name")
	shardIteratorTypePtr := flag.String("i", "LATEST", "Shard iterator type.")
	regionPtr := flag.String("r", "", "AWS region, e.g. 'us-west-2'")
	flag.Parse()
	awsKinesisStreamName := *streamPtr
	awsKinesisShardIteratorType := *shardIteratorTypePtr
	awsRegion := *regionPtr

	shards, err := kinchan.GetShards(awsKinesisStreamName, awsRegion)
	if err != nil {
		log.Fatal(err)
	}

	messageChannel := make(chan []byte, 1000)
	for _, shard := range shards {
		go kinchan.Consume(shard, awsKinesisStreamName, awsKinesisShardIteratorType, messageChannel)
	}

	go logMessagesToConsole(messageChannel)

	waitForever := make(chan string)
	waitForever <- ""
}

func logMessagesToConsole(messageChannel chan []byte) {
	for {
		data := <-messageChannel
		fmt.Println(string(data))
	}
}```

## Caveats

* Consumer only.
* No reliable consumer implementation.
* No automatic resharding.

## TODO

* Better error handling
* Producer
* Test automation
