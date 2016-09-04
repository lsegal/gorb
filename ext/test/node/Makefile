
CGO_CFLAGS = -I $(shell ruby -rrbconfig -e 'puts RbConfig::CONFIG["rubyhdrdir"]') -I $(shell ruby -rrbconfig -e 'puts RbConfig::CONFIG["rubyarchhdrdir"]')
CGO_LDFLAGS = $(shell ruby -rrbconfig -e 'puts RbConfig::CONFIG["LIBRUBYARG"]')
EXT = $(shell ruby -rrbconfig -e 'puts RbConfig::CONFIG["DLEXT"]')
export CGO_CFLAGS
export CGO_LDFLAGS

all:
	go build -buildmode=c-shared -o node.${EXT} .
