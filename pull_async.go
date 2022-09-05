// [START pubsub_subscriber_async_pull]
// [START pubsub_quickstart_subscriber]
package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync/atomic"

	// "time"

	"cloud.google.com/go/pubsub"
)

func pullMsgs(w io.Writer, projectID, subID string) error {
	// projectID := "my-project-id"
	// subID := "my-sub"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	sub := client.Subscription(subID)

	// Receive messages for 10 seconds, which simplifies testing.
	// Comment this out in production, since `Receive` should
	// be used as a long running operation.
	// ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	// defer cancel()

	var received int32
	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		fmt.Fprintf(w, "Got message: %q\n", string(msg.Data))
		atomic.AddInt32(&received, 1)
		msg.Ack()
	})
	if err != nil {
		return fmt.Errorf("sub.Receive: %v", err)
	}
	fmt.Fprintf(w, "Received %d messages\n", received)

	return nil
}

func main() {
	// w := io.Writer(os.Stdout)
	pullMsgs(os.Stdout, "cliu201", "my-topic-sub")

}

// [END pubsub_subscriber_async_pull]
