CGO_CFLAGS = -I $(shell ruby -rrbconfig -e 'puts RbConfig::CONFIG["rubyhdrdir"]') -I $(shell ruby -rrbconfig -e 'puts RbConfig::CONFIG["rubyarchhdrdir"]')
CGO_LDFLAGS = $(shell ruby -rrbconfig -e 'puts RbConfig::CONFIG["LIBRUBYARG"]')
export CGO_CFLAGS
export CGO_LDFLAGS

all:
	go build .

test:
	bash test/test.sh

.PHONY: all test clean
