BINARY=myGoSamples
VERSION=1.0.0
BUILD=`git rev-parse HEAD`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

build:
	go build ${LDFLAGS} -o ${BINARY}

install:
	go install ${LDFLAGS}

clean:
	if [ -f ${BINARY} ] ; then rm  ${BINARY} ; fi

.PHONY: clean install
