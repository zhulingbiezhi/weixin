FROM zhulingbiezhi/weixin:master

MAINTAINER ququ "key.hu@bindo.com"

WORKDIR $GOPATH/src/weixin
ADD . $GOPATH/src/weixin

RUN source ~/.bash_profile
RUN go build 
RUN /etc/init.d/mysql start

EXPOSE 9091

ENTRYPOINT ["./weixin"]