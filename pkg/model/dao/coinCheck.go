package dao

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"bit-board/pkg/model/dto"
	"strconv"
	"time"
)

var url = "https://coincheck.com/api/rate/btc_jpy"

func getNowJPY(buyBTC float64)(float64,error){
	var response dto.GetNowJPY
	var nowJPY float64
	var err error
	resp, _ := http.Get("https://coincheck.com/api/exchange/orders/rate/?order_type=sell&pair=btc_jpy&amount="+strconv.FormatFloat(buyBTC, 'f', 4, 64))
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(byteArray, &response)
	if err !=nil{
		log.Println(err)
		return nowJPY, err
	}
	nowJPY,err=strconv.ParseFloat(response.Price, 64)
	if err !=nil{
		log.Println(err)
		return nowJPY, err
	}
	return nowJPY,err
}

func GetNowRate()(dto.BitcoinRate,error){
	var bitCoin dto.BitcoinRate
	var coinCheck dto.CoinCheck
	var err error
	bitCoin.Timestamp=int(time.Now().Unix())
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(byteArray, &coinCheck)
	if err !=nil{
		log.Println(err)
		return bitCoin, err
	}
	bitCoin.NowRate,err=strconv.ParseFloat(coinCheck.NowRate, 64)
	if err !=nil{
		log.Println(err)
		return bitCoin, err
	}
	return bitCoin,err
}
