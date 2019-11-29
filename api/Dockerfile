FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY . /go/src/github.com/user/GoNeo/api
WORKDIR /go/src/github.com/user/GoNeo/api

RUN go get ./vendor/database

RUN go get ./
RUN go build

CMD if [ ${APP_ENV} = production ]; \
	then \
	api; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi
	
EXPOSE 8080
