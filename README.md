```
go mod init github.com/liuchenggang2014/pubsub_pull_msg

go get cloud.google.com/go/pubsub

go mod tidy

go run pull_async.go

# publish mesage into topic
gcloud pubsub topics publish my-topic \
    --message=hello_zulong \
    --attribute=foo="bar"
```