FROM golang:latest  
  
MAINTAINER huhai <zhulingbiezhi@gmail.com>  
RUN echo $GOPATH
WORKDIR $GOPATH/bin

RUN cd $GOPATH; \
	mkdir src; \
	mkdir bin; \  
    cd src; \  
    git clone https://github.com/zhulingbiezhi/weixin.git -b $ENV; \ 
	cd weixin; \
	go build -o ../bin
	
EXPOSE 8080

ENTRYPOINT ["./weixin"]
