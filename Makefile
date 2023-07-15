proto: gochess/position.proto
	protoc -I=gochess/ gochess/position.proto --go_out=./

run: proto
	go run main.go

clean:
	rm -r */*.pb.go