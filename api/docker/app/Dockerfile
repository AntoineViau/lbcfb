FROM golang:latest 
RUN mkdir /lbcfb
RUN git clone https://github.com/AntoineViau/lbcfb.git /lbcfb
WORKDIR /lbcfb/src/fizzbuzz
RUN go test
WORKDIR /lbcfb/src
RUN go build 
EXPOSE 8080
CMD "/lbcfb/src/lbcfb"
