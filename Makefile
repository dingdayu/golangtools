install:
	go install ./example/config/

test: install
	go run ./example/config/main.go -c ./example/config/conf.json
	config -c ./example/config/conf.toml


fmt:
	gofmt -w *.go */*.go
	colcheck *.go */*.go

tags:
	find ./ -name '*.go' -print0 | xargs -0 gotags > TAGS

push:
	git push origin master
	git push github master
