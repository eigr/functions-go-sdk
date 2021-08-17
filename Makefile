.PHONY: test tck compile-protobuf

test:
	go test -v -race ./...

tck:
	./tck/run_tck.sh

compile-protobuf:
	./build/compile-pb.sh