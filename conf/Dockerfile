FROM golang:latest  
  
MAINTAINER huhai <zhulingbiezhi@gmail.com>  
RUN echo $GOPATH
WORKDIR $GOPATH/src/weixin

RUN cd $GOPATH; \
	mkdir src; \
	mkdir bin; \  
    cd src; \  
    git clone https://github.com/zhulingbiezhi/weixin.git -b master; \ 
	cd weixin; \
	go build 
	
EXPOSE 8080

ENTRYPOINT ["./weixin"]