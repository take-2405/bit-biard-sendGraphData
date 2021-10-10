package main

import (
	"bit-board/pkg/model/dao"
	"bit-board/pkg/model/dto"
	"bit-board/pkg/view"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func handler() (events.APIGatewayProxyResponse,error) {
	var response dto.Response
	client, err := dao.New()
	if err != nil {
		log.Println(err)
		return view.ReturnInternalServerErrorResponse(err)
	}

	Graph,err:=client.BitBorad.BitBoardLogic.GetGraphData()
	if err != nil {
		log.Println(err)
		return view.ReturnInternalServerErrorResponse(err)
	}
	for i := len(Graph)-1;i>=0;i--{
		response.Rate=append(response.Rate,Graph[i].Rate)
		response.Label=append(response.Label,Graph[i].Label)
	}


	nowRate,err:= dao.GetNowRate()
	if err != nil {
		log.Println(err)
		return view.ReturnInternalServerErrorResponse(err)
	}

	response.Timestamp=nowRate.Timestamp
	response.NowRate=nowRate.NowRate

	resJSON, _ := json.Marshal(response)

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "origin,Accept,Authorization,Content-Type",
			"Content-Type":                 "application/json",
		},
		Body:       string(resJSON),
		StatusCode: 200,
	},nil
}

func main(){
	lambda.Start(handler)
}
