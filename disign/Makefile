SOURCES=$(wildcard *.go)
EXEC=disign
BINARIES=../bin
all: $(EXEC)

$(EXEC):
	go build -o $(BINARIES)/$(EXEC)

clean:
	$(RM) $(BINARIES)/$(EXEC)
	
dependencies:
	go get -u github.com/spf13/cobra/cobra
	go get github.com/spf13/viper
	go get github.com/mitchellh/go-homedir
