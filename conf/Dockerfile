FROM golang:latest  
  
MAINTAINER huhai <zhulingbiezhi@gmail.com>  

WORKDIR $GOPATH/src/weixin

RUN cd $GOPATH; \
	mkdir src; \
	mkdir bin; \  
    cd src; \  
    git clone https://github.com/zhulingbiezhi/weixin.git -b dev; \ 
	cd weixin; \
	go build 
	
EXPOSE 8080

ENTRYPOINT ["./weixin"]