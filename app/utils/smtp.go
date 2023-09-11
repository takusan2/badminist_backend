package utils

// func SendEmail(subject string, body string, recipients []string) {
// 	auth := smtp.CRAMMD5Auth(username, password)
// 	msg := []byte(strings.ReplaceAll(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", strings.Join(recipients, ","), subject, body), "\n", "\r\n"))
// 	if err := smtp.SendMail(fmt.Sprintf("%s:%d", hostname, port), auth, from, recipients, msg); err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 	}
// }
