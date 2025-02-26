package publisher

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"meme_service/internal/app/meme/list/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	snstypes "github.com/aws/aws-sdk-go-v2/service/sns/types"
)

var topicArn = "arn:aws:sns:us-east-1:000000000000:token_balance_topic"

type broker struct {
  snsClient *sns.Client
}

func (s broker) Publish(input types.InputPublisher) error {
  var buf bytes.Buffer
  json.NewEncoder(&buf).Encode(input)

  message := &sns.PublishInput{
    Message:  aws.String(buf.String()),
    TopicArn: aws.String(topicArn),
    MessageAttributes: map[string]snstypes.MessageAttributeValue{
      "customer_id": {
        DataType:    aws.String("String"),
        StringValue: aws.String(input.GetCustomerID()),
      },
    },
  }

  _, err := s.snsClient.Publish(context.TODO(), message)
  if err != nil {
    return err
  }

  fmt.Println("Mensagem publicada no topico com sucesso")
  return nil
}

func New(snsClient *sns.Client) types.Publisher {
  return broker{
    snsClient: snsClient,
  }
}

