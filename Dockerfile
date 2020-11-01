#FROM  telkomindonesia/alpine:go-1.13

#LABEL maintainer="adisaputra.id@gmail.com"

#WORKDIR /usr/src/app

# Update package
#RUN apk add --update --no-cache --virtual .build-dev build-base python python-dev git

#COPY . .

#RUN make install \
 # && make build

# Expose port
#EXPOSE 9000

# Run application
#CMD ["make", "start"]


FROM  telkomindonesia/debian:php-7.2-nginx-novol
copy . .
