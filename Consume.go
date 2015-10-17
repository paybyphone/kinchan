package kinchan

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

//Consume gets data from a shard of an AWS Kinesis stream and puts them in a Go channel, once every second.
func Consume(shardID string, streamName string, shardIteratorType string, eventChannel chan []byte) {
	svc := kinesis.New(nil)
	getShardIteratorParams := &kinesis.GetShardIteratorInput{
		ShardId:           aws.String(shardID),           // Required
		ShardIteratorType: aws.String(shardIteratorType), // Required
		StreamName:        aws.String(streamName),        // Required
	}
	shardIterator, err := svc.GetShardIterator(getShardIteratorParams)
	if err != nil {
		log.Panic(err)
	}

	nextShardIterator := shardIterator.ShardIterator
	for {
		getRecordsParams := &kinesis.GetRecordsInput{
			ShardIterator: nextShardIterator, // Required
			Limit:         aws.Int64(1000),
		}
		getRecordsResp, err := svc.GetRecords(getRecordsParams)
		if err != nil {
			log.Panic(err)
		}
		for _, record := range getRecordsResp.Records {
			eventChannel <- record.Data
		}
		nextShardIterator = getRecordsResp.NextShardIterator
		time.Sleep(time.Second)
	}
}
