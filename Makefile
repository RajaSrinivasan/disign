SOURCES=$(wildcard *.go)
EXEC=disign
BINARIES=../../bin
all: $(EXEC)

$(EXEC): 
	go build -o $(BINARIES)/$(EXEC)

version:
	cd cmd; ../$(BINARIES)/srctrace --language go -m 0 -n 1 -b 1 --output cmd

clean:
	$(RM) $(BINARIES)/$(EXEC)
	
dependencies:
	go get -u github.com/spf13/cobra/cobra
	go get github.com/spf13/viper
	go get github.com/mitchellh/go-homedir
	go get -u -v golang.org/x/crypto/ssh
	go get -u -v github.com/ianmcmahon/encoding_ssh
