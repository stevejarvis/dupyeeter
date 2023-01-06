module dupyeeter

go 1.19

require (
	github.com/aws/aws-lambda-go v1.36.0
	github.com/joeshaw/envdecode v0.0.0-20200121155833-099f1fc765bd
	github.com/stevejarvis/dupyeeter v0.0.0-00010101000000-000000000000
	golang.org/x/oauth2 v0.4.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.5.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace github.com/stevejarvis/dupyeeter => ./
