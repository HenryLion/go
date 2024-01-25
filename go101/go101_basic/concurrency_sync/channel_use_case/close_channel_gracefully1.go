package main

// One general principle of using Go channels is
// don't close a channel from the receiver side
// and don't close a channel if the channel has
// multiple concurrent senders.
// In other words, we should only close a channel
// in a sender goroutine if the sender is the only sender of the channel.