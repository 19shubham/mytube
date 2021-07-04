FROM golang:1.15.7-buster

EXPOSE 80

WORKDIR /usr/local/mytube

ENV GOPATH=/usr/local/mytube/vendor:/usr/local/mytube \
    PATH=/usr/local/go/bin:${GOPATH}:${PATH}

RUN rm -f /etc/localtime && \
 ln -s /usr/share/zoneinfo/Asia/Kolkata /etc/localtime

COPY ./ /usr/local/mytube

RUN mkdir -p bin && \
	make default
