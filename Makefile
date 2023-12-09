all:
	go install app
	go install utils
clean:
	rm ./bin/*.exe
	RM ./pkg/windows_amd64/*.a