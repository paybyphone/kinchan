package kinchan

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

//GetShards returns shard names for an AWS Kinesis stream.
func GetShards(streamName string, region string) ([]string, error) {
	defaults.DefaultConfig.Region = aws.String(region)
	svc := kinesis.New(nil)

	params := &kinesis.DescribeStreamInput{
		StreamName: aws.String(streamName),
	}
	resp, err := svc.DescribeStream(params)
	if err != nil {
		return nil, err
	}

	var shards []string
	for _, shard := range resp.StreamDescription.Shards {
		shards = append(shards, *shard.ShardId)
	}
	return shards, nil
}
