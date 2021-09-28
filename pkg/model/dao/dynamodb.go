package dao

import (
	"bit-board/pkg/model/dto"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type  DynamoDB struct{
	Dynamo  *dynamodb.DynamoDB
	BitBorad Methods
}

type Methods struct {
	BitBoardLogic methods
}

type methods interface {
	GetGraphData( )([]dto.Graph,error)
}

func New()(*DynamoDB,error){
	//DB接続
	svc := dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))

	// init methods
	bitboardtMethod := newBitBoardClient(svc)

	return &DynamoDB{
		Dynamo:  svc,
		BitBorad: Methods{bitboardtMethod},
	},nil
}