# See README.txt.

.PHONY: all go clean

all: go

go: *.proto
	#protoc --go_out=plugins=$$GOPATH/src ./charger.proto
	#ls -1 *.proto | while read file ;do protoc --go_out=$$GOProjectPath/src ./$$file; done
	protoc --go_out=plugins=grpc:`PWD`/../../../../../  ./charger.proto
	protoc --go_out=plugins=grpc:`PWD`/../../../../../  ./measure.proto
	protoc --go_out=plugins=grpc:`PWD`/../../../../../  ./warning.proto
	# protoc $$PROTO_PATH --cpp_out=. --java_out=. --python_out=. addressbook.proto
clean:
	rm -rf $$GOProjectPath/src/github.com/Kotodian/protocol/golang/hardware/charger
# go: addressbook.proto
