package jenkins

// Client --
type Client interface {
	DeleteJob() error
}

// JClient -- jenkins clients
type JClient struct {
}
