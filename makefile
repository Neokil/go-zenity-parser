install:
	go build -o go-zenity-parser
	chmod a+x go-zenity-parser
	ln -sf $(shell pwd)/go-zenity-parser /usr/local/bin/go-zenity-parser
