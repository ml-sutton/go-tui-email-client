package lifecycle

type Channels struct {
}

func CreateChannels() *Channels {
	return &Channels{}
}

func (c *Channels) Close() {

}
