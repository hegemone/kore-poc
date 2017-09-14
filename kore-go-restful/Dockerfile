FROM golang


COPY ./ /usr/local/go/src/github.com/dahendel/kore-poc
WORKDIR /usr/local/go/src/github.com/dahendel/kore-poc
RUN apt-get update && apt-get install golang-glide git gcc -y

RUN glide install && \
 mkdir 3rdparty && cd 3rdparty && \
 git clone -b 2.x https://github.com/swagger-api/swagger-ui.git && \
 cd ../docs && ln -s ../3rdparty/swagger-ui/dist swagger-ui && cd ../ && \
 go build

EXPOSE 8080
CMD ./kore-poc
