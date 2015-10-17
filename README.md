# AWS Kinesis Channels

Makes Kinesis streams look like Go channels.

## Usage:

```
package main

import (
	"fmt"
	"log"

	"github.com/paybyphone/kinchan"
)

func main() {
	awsKinesisStreamName := "MyStreamName"
	awsKinesisShardIteratorType := "LATEST"
	awsRegion := "us-west-2"

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
}
```

## Caveats

* Consumer only.
* No reliable consumer implementation.
* No automatic resharding.
