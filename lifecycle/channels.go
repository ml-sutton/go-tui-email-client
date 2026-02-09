package lifecycle

type Channels struct {
	fatal chan error
}

func CreateChannels() *Channels {
	return &Channels{
		fatal: make(chan error, 1),
	}
}

func (c *Channels) Close() {
	close(c.fatal)
}
