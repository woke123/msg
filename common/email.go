package common

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"gopkg.in/gomail.v2"
	"msg/config"
)

func InitEmail(redis_c redis.Conn){
	startEmail(redis_c,config.C.Msg.Email_sent)
	return
}

func startEmail(redis_c redis.Conn,email_key string){
	for {
		datas, err := redis.StringMap(redis_c.Do("brpop", email_key,60))
		if err != nil || datas != nil {
			result := &config.EmailData{}
			err = json.Unmarshal([]byte(datas[email_key]), &result)
			if err != nil {
				continue
			}
			go sendEmail(result.Email, result.Content, result.Subject, redis_c)
		}
	}
}

func sendEmail(email string, content string,subject string,redis_c redis.Conn){
	m := gomail.NewMessage()
	m.SetHeader("From",config.C.Email.User)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)
	d := gomail.NewDialer(config.C.Email.Host, config.C.Email.Port, config.C.Email.User, config.C.Email.Pass)
	err := d.DialAndSend(m)
	if err != nil {
		fail := make(map[string]string, 2)
		fail["email"] = email
		fail["content"] = content
		fail["subject"] = subject
		mjson,_ :=json.Marshal(fail)
		fail_data :=string(mjson)
		redis_c.Do("lpush", config.C.Msg.Email_fail,fail_data)
	}
	return
}