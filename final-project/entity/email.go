package entity

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/kriti-mauludin/final-project/internal/mailer"
)

type SendEmailReq struct {
	UserID    int64     `json:"user_id"`
	TypeEmail string    `json:"type_email"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Datetime  string    `json:"datetime"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *SendEmailReq) LoadFromMap(m map[string]interface{}) error {
	data, err := json.Marshal(m)
	if err == nil {
		err = json.Unmarshal(data, c)
	}
	return err
}

func FillMailObjLogin(data SendEmailReq) (mailer.MailObject, error) {
	mailObj := mailer.MailObject{}
	switch data.TypeEmail {
	case "success-login":
		mailObj.Subject = "Notifikasi Login Berhasil"
	}
	buf, errGetWD := getTemplate(data.TypeEmail, data)
	if errGetWD != nil {
		return mailObj, errGetWD
	}
	mailObj.Email = data.Email
	mailObj.ParsedBody = buf.String()
	return mailObj, errGetWD
}

func getTemplate(typeEmail string, tempEmail interface{}) (*bytes.Buffer, error) {
	var buf *bytes.Buffer
	wd, errGetWD := os.Getwd()
	if errGetWD != nil {
		log.Println(errGetWD)
		return buf, errGetWD
	}
	pathHTML := wd + "/internal/mailer/template/" + typeEmail + ".html"
	t, err := template.ParseFiles(pathHTML)
	if err != nil {
		return buf, err
	}
	buf = new(bytes.Buffer)
	if err = t.Execute(buf, tempEmail); err != nil {
		return buf, err
	}
	return buf, nil
}
