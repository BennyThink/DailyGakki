// DailyGakki - config
// 2020-10-17 14:14
// Benny <benny.think@gmail.com>

package main

var photos = "https://photos.app.goo.gl/2aLeoBiRypWRR8yY9"

type Database struct {
	ChatId string `json:"chat_id"`
	Count  string
	Time   int64
}
