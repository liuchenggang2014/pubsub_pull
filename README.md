```
go mod init github.com/liuchenggang2014/pubsub_pull_msg

go mod tidy

go run pull_async.go

# publish mesage into topic
gcloud pubsub topics publish zulong \
    --message=hello_zulong \
    --attribute=foo="bar"
```