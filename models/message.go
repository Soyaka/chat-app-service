package models



func NewMessage(content, from, to, typee string) *message {
	return &message{
		Content: content,
		From: from,
		To: to,
		Type: typee,
	}
}

type message struct{
	Content string
	From string
	To string
	Type string

}

