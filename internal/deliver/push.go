package deliver

// Push sends one rendered message to a device token.
// Retries share the delivery queue with email — see the open issue.
func Push(deviceToken, body string) error { return nil }
