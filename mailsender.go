package main

import (
  "io"
  "io/ioutil"
  "os"
  "log"
  "fmt"
  "time"
  "net/smtp"
  "strings"
  "encoding/csv"
)

func main() {
	// parameters
	from := os.Getenv("FROM")
	from_name := os.Getenv("FROM_NAME")
	subj := os.Getenv("SUBJ")
	dryRun := os.Getenv("DRYRUN")

	// smtp server configuration.
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	// Args
	recipients_file := os.Args[1]
	letter_file := os.Args[2]

	// Read letter
	letter, err := ioutil.ReadFile(letter_file)
	if err != nil {
		log.Fatal("Error while reading the body file", err)
		return
	    }

	// CSV reader
	recipients, err := os.Open(recipients_file)
	if err != nil {
		log.Fatal("Error while reading the recipient list file", err)
		return
	    }

	r := csv.NewReader(recipients)

	defer recipients.Close()

	// Authentication.
	auth := smtp.PlainAuth(
		"",
		smtpUser,
		smtpPass,
		smtpHost)

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		to := record[0]
		first := record[1]
		last := record[2]
		full := first + " " + last
		to_hdr := first + " " + last + " <" + to + ">"
		body := strings.Replace(
			string(letter),
			"%RECIPIENT%", full, -1)

		msg := "From: " + from_name + " <" + from + ">\n" +
			"To: " + to_hdr + "\n" +
			"Subject: " + subj + "\n\n" +
			body

		// Sending email.
		if (len(dryRun) > 0) {
			fmt.Println("DRY RUN: " + from + " -> " + to)
		} else {
			err = smtp.SendMail(
				smtpHost + ":" + smtpPort,
				auth,
				from,
				[]string{to},
				[]byte(msg))

			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println(from + " -> " + to)
			}
		}

		time.Sleep(1 * time.Second)
	}
}
