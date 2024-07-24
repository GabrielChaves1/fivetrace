package ports

type Emailer interface {
	SendHTMLEmail(to, subject, htmlBody string) error
}