RUN cd /tmp && \
 wget https://storage.googleapis.com/golang/go1.14.linux-amd64.tar.gz && \
 tar -C /usr/local -xzf go1.14.linux-amd64.tar.gz

EXPOSE 80

WORKDIR /usr/local/mytube

ENV GOPATH=/usr/local/goibibo/mytube/vendor:/usr/local/mytube \
    PATH=/usr/local/go/bin:${GOPATH}:${PATH}

RUN rm -f /etc/localtime && \
 ln -s /usr/share/zoneinfo/Asia/Kolkata /etc/localtime

COPY ./ /usr/local/mytube

RUN mkdir -p bin && \
	make default
