SHELL = /bin/sh
HERE = $(shell pwd)
BIN = $(HERE)/bin
GPM = $(HERE)/gpm
DEPS = $(HERE)/.godeps
GOPATH := $(DEPS):$(HERE)
GOBIN = $(BIN)

PLATFORM=$(shell uname)

# Setup commands and env vars if there is no system go linked into bin/go
PATH := $(HERE)/bin:$(DEPS)/bin:$(PATH)

PACKAGE = github.com/mozilla-services/pushgo

.PHONY: all build clean test simplepush memcached

all: build

$(BIN):
	mkdir -p $(BIN)

$(DEPS):
	@echo "Installing dependencies"
	$(GPM) install

build: $(DEPS)

libmemcached-1.0.18:
	wget -qO - https://launchpad.net/libmemcached/1.0/1.0.18/+download/libmemcached-1.0.18.tar.gz | tar xvz
	cd libmemcached-1.0.18 && \
	./configure --prefix=/usr && \
	autoreconf -ivf
ifeq ($(PLATFORM),Darwin)
	cd libmemcached-1.0.18 && \
	sed -i '' $$'/ax_pthread_flags="pthreads none -Kthread -kthread lthread -pthread -pthreads -mthreads pthread --thread-safe -mt pthread-config"/c\\\nax_pthread_flags=\"pthreads none -Kthread -kthread lthread -lpthread -lpthreads -mthreads pthread --thread-safe -mt pthread-config"\n' m4/ax_pthread.m4
else
	cd libmemcached-1.0.18 && \
	sed -i '/ax_pthread_flags="pthreads none -Kthread -kthread lthread -pthread -pthreads -mthreads pthread --thread-safe -mt pthread-config"/c\ax_pthread_flags="pthreads none -Kthread -kthread lthread -lpthread -lpthreads -mthreads pthread --thread-safe -mt pthread-config"' m4/ax_pthread.m4
endif

memcached: libmemcached-1.0.18
	cd libmemcached-1.0.18 && sudo make install

simplepush:
	rm -f simplepush
	@echo "Building simplepush"
	go build -o simplepush github.com/mozilla-services/pushgo

test:
	go test $(addprefix $(PACKAGE)/,id simplepush)

vet:
	go vet $(addprefix $(PACKAGE)/,client id simplepush)

clean:
	rm -rf bin $(DEPS)
	rm -f simplepush