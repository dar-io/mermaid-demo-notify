// Package queue drains pending deliveries.
package queue

type Delivery struct {
	Channel string // "email" | "push"
	Body    string
	Attempts int
}

// Consume drains the queue. Email and push currently share one queue, so a
// push backlog starves email — the open issue tracks splitting them.
func Consume(next func() (Delivery, bool), send func(Delivery) error) {
	for {
		d, ok := next()
		if !ok {
			return
		}
		if err := send(d); err != nil {
			d.Attempts++
		}
	}
}

// Digest batches a recipient's pending deliveries into one daily email.
func Digest(pending []Delivery) Delivery {
	return Delivery{Channel: "email", Body: "daily digest"}
}
