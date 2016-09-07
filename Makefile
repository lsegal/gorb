CGO_CFLAGS = -I $(shell ruby -rrbconfig -e 'puts RbConfig::CONFIG["rubyhdrdir"]') -I $(shell ruby -rrbconfig -e 'puts RbConfig::CONFIG["rubyarchhdrdir"]')
CGO_LDFLAGS = -L $(shell ruby -rrbconfig -e 'puts RbConfig::CONFIG["libdir"]') -l$(shell ruby -rrbconfig -e 'puts RbConfig::CONFIG["RUBY_SO_NAME"]')
export CGO_CFLAGS
export CGO_LDFLAGS

all:
	go build .

test:
	bash test/test.sh

.PHONY: all test clean
