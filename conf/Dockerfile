FROM zhulingbiezhi/weixin:master

MAINTAINER ququ "key.hu@bindo.com"

WORKDIR $GOPATH/src/weixin
ADD . $GOPATH/src/weixin
RUN go build 

EXPOSE 9091

ENTRYPOINT ["./weixin"]