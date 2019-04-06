package mocks

type Queue struct{}

// Peek returns a message from the queue but does not delete it.
func (q *Queue) Peek() (string, error) {
	return "", nil
}

// Delete deletes the given message.
func (q *Queue) Delete(msg string) error {
	return nil
}
