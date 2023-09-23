# Simple Mail Sender

A simple mail sender that takes recipients from a CSV file.

```
$ go build

$ export DRYRUN=true
$ export SMTP_USER=from@example.com
$ export SMTP_PASS=your_password
$ export FROM_NAME="Your Name"
$ export FROM="from@example.com"
$ export SMTP_HOST="smtp.gmail.com"
$ export SMTP_PORT=587
$ export SUBJ="Here's a letter for you"

$ ./mailsender email-first-last.csv letter.txt
from@example.com -> to@address.com
```

Remember to un-set the `DRYRUN` variable when you really want to send out.
