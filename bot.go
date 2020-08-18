package main

import "fmt"
import "encoding/json"
import "io/ioutil"
//import "bytes"
import "net/http"
import "strconv"


const telegramBaseUrl="https://api.telegram.org/bot"
const telegramToken="xxx"
const methodGetMe="getMe"
const methodGetUpdates="getUpdates"
const methodSendMessage="sendMessage"
func main(){
	

	/*
	err :=json.Unmarshal(test, &getMe)
	if err!=nil{
		fmt.Println(err.Error())
	}

	//fmt.Printf("%v",getMe)
	fmt.Println(getMe.Result.UserName)*/
	//fmt.Println(getUrlByMethod(methodGetMe))

	body:=getBodyUrlAndData(getUrlByMethod(methodGetUpdates))
//	fmt.Println(string(body))
	//fmt.Printf("%s",body)
//	getMe:= GetMeT{}
	getUpdates:=GetUpdatesT{}
	json.Unmarshal(body,&getUpdates)
	//fmt.Printf("%v",getUpdates)
	
	sendMessageUrl := getUrlByMethod(methodSendMessage)
	for _,item := range(getUpdates.Result){
		if item.Message.Text == "test"{
			chatId:=strconv.Itoa(item.Message.Chat.ID)
			targetUrl :=sendMessageUrl + "?chat_id=" + chatId + "&text=" + item.Message.Text
			
			getBodyUrlAndData(targetUrl)
			
		}
		
	}

}
func getUrlByMethod(methodName string) string{
	return telegramBaseUrl + telegramToken + "/" + methodName
}
func getBodyUrlAndData(url string)[]byte{
	//r:=bytes.NewReader(data)
	resp,err:=http.Get(url)
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)

	if err!=nil{
		fmt.Println(err.Error())
	}
	return body
}
type GetMeT struct{
	Ok bool `json:"ok"`
	Result GetMeResultT `json:"result"`

}

type GetMeResultT struct{
	Id int `json:"id"`
	IsBot bool `json:"is_bot"`
	FirstName string `json:"first_name"`
	UserName string `json:"username"`
}

type GetUpdatesT struct{
	Ok bool `json:"ok"`
	Result []GetUpdatesResultT `json:"result"`
}

type SendMessageT struct{
	Ok bool `json:"ok"`
	Result MessageT `json:"result"`
}

type MessageT struct{
	MessageID int `json:"message_id"`
	From GetUpdatesResultMessageFromT `json:"from"`
	Chat GetUpdatesResultMessageChatT `json:"chat"`
	Date int `json:"date"`
	Text string `json:"text"`
}

type GetUpdatesResultT struct{
	UpdateID int `json:"update_id"`
	Message MessageT `json:"message,omitempty"`
}
type GetUpdatesResultMessageFromT struct{
	ID int `json:"id"`
	IsBot bool `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username string `json:"username"`
	LanguageCode string `json:"language_code"`
}
type GetUpdatesResultMessageChatT struct{
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	Username string `json:"username"`
	Type string `json:"type"`
}