package queue

type Queue struct{}

// Peek returns a message from the queue but does not delete it.
func (q *Queue) Peek() (string, error) {
	return "", nil
}

// Insert inserts a single message to the queue.
func (q *Queue) Insert(msg string) error {
	return nil
}

// InsertBatch inserts a batch of messages into the queue.
func (q *Queue) InsertBatch(msg []string) error {
	return nil
}

// Delete deletes the given message.
func (q *Queue) Delete(msg string) error {
	return nil
}

// Pop pulls a message from the queue, deletes it and returns it.
func (q *Queue) Pop() (string, error) {
	return "", nil
}
