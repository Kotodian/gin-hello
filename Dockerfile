# Version 0.1

# base image
FROM golang:1.16-alpine

# maintainer
MAINTAINER 417350372@qq.com

# expose port
EXPOSE 9091

# container
CMD ["./gin-hello"]



