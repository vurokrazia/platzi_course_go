FROM golang:1.15.0

ENV APP_NAME codigo
ENV PORT 8080

COPY . /go/src/ 
#${APP_NAME}
WORKDIR /go/src/  
#${APP_NAME}

#RUN go get ./
#RUN go build -o ${APP_NAME}

#CMD ./${APP_NAME}

#EXPOSE ${PORT}

CMD ["bash"] 