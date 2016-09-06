
CGO_CFLAGS = -I $(shell ruby -rrbconfig -e 'puts RbConfig::CONFIG["rubyhdrdir"]') -I $(shell ruby -rrbconfig -e 'puts RbConfig::CONFIG["rubyarchhdrdir"]')
CGO_LDFLAGS = -L $(shell ruby -rrbconfig -e 'puts RbConfig::CONFIG["libdir"]') -l$(shell ruby -rrbconfig -e 'puts RbConfig::CONFIG["RUBY_SO_NAME"]')
EXT = $(shell ruby -rrbconfig -e 'puts RbConfig::CONFIG["DLEXT"]')
export CGO_CFLAGS
export CGO_LDFLAGS

all:
	go build -buildmode=c-shared -o data.${EXT} .
