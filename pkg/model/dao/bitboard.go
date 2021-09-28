package dao

import (
	"bit-board/pkg/model/dto"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"math"
	"strconv"
	"time"
)

type InsertData struct {
	UserId  string  `dynamodbav:"userID"`
	BuyTime int  `dynamodbav:"buyTime"`
	BuyJPY  int `dynamodbav:"buyJPY"`
	BuyBTC float64 `dynamodbav:"buyBTC"`
}

type bitBoardMethods struct {
	Client *dynamodb.DynamoDB
}

func newBitBoardClient(client *dynamodb.DynamoDB) methods {
	return &bitBoardMethods{Client: client}
}

func (r *bitBoardMethods)GetGraphData()([]dto.Graph,error){
	var err error
	var graph []dto.Graph
	var item dto.Graph
	getParam := &dynamodb.QueryInput{
		TableName: aws.String("GraphBTC"),
		ExpressionAttributeNames: map[string]*string{
			"#ID":       aws.String("ID"),
			"#timestamp": aws.String("timestamp"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":ID": {
				N: aws.String("1"),
			},
			":timestamp": {
				N: aws.String(strconv.FormatInt(time.Now().Unix(), 10)),
			},
		},
		ScanIndexForward:       aws.Bool(false), // ソートキーのソート順（指定しないと昇順）
		KeyConditionExpression: aws.String("#ID = :ID AND #timestamp <= :timestamp"), // 検索条件
		Limit: aws.Int64(20),
	}

	results, err := r.Client.Query(getParam)
	if err != nil {
		log.Println(err)
		return 	graph,err
	}

	for _,j := range results.Items {
		err = dynamodbattribute.UnmarshalMap(j, &item)
		if err != nil {
			log.Println(err)
			return 	graph,err
		}
		graph = append(graph,item)
	}

	return graph,err
}