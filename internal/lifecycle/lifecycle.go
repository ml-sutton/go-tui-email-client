package lifecycle

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func RunClient() error {
	var ctx context.Context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	var channels *Channels = CreateChannels()
	defer cancel()
	defer channels.Close()

	select {
	case <-ctx.Done():
		return nil
	case err, ok := <-channels.fatal:
		if !ok {
			return fmt.Errorf("hello")
		}
		return err
	}
}
