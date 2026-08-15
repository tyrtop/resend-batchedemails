package main

import (
	"bytes"
	"fmt"
	_ "embed"
	"text/template"
	"log"
	"github.com/resend/resend-go/v3"
)

//go:embed embed.html
var credentialsHTML string

type emailData struct {
	User 
	JellyfinURL string
	JellyfinHost string 
	JellyseerrURL string
	FromAddress string

}

func main() {
	hostURL := "https://watch.tyrtop.com"
	jellyfin := "watch.tyrtop.com"
	jellyseerrURL := "https://requests.tyrtop.com"
	fromAddress := "noreply@tyrtop.com"
	
	users, err := loadUsers("users.json")
	if err != nil {
		log.Fatal(err)
	}

	tmpl := template.Must(template.New("cred").Parse(credentialsHTML))

	var batch = []*resend.SendEmailRequest{}
	
	for _, u := range users {
		fmt.Printf("%+v\n", users)
		var buf bytes.Buffer
		d := emailData{u, hostURL, jellyfin, jellyseerrURL, fromAddress} 
		if err:= tmpl.Execute(&buf, d); err != nil {
			log.Fatal(err)
		}
		batch = append(batch, &resend.SendEmailRequest{
			From: fromAddress,
			To: []string{u.Email},
			Subject: "Account Credentials",
			Html: buf.String(),
		})
	}

	client := resend.NewClient("re_***")

	sent, err := client.Batch.Send(batch)
	if err != nil{
		panic(err)
	}
	fmt.Println(sent.Data)
}
