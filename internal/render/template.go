// Package render turns an event into a message body.
package render

import "text/template"

var receipt = template.Must(template.New("receipt").Parse(
	"Your payment of {{.Amount}} was received."))

func Receipt(amount string) (string, error) {
	return execute(receipt, map[string]string{"Amount": amount})
}

func execute(t *template.Template, data any) (string, error) { return "", nil }
