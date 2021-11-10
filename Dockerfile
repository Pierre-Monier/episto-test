# The base go-image
FROM golang:1.17-alpine

# Create a directory for the app
RUN mkdir /app

# Copy all files from the current directory to the app directory
COPY . /app

# Set working directory
WORKDIR /app

# Run command as described:
# go build will build an executable file named server in the current directory
RUN go build -o server .

# Run the server executable
CMD [ "/app/server" ]

# This can be really cool for dev, but can't make it work with docker volume :(
# FROM golang

# ARG app_env
# ENV APP_ENV $app_env

# COPY . /go/src/github.com/user/myProject/app
# WORKDIR /go/src/github.com/user/myProject/app

# RUN go get ./
# RUN go build

# CMD if [ ${APP_ENV} = production ]; \
#     then \
#     app; \
#     else \
#     go get github.com/pilu/fresh && \
#     fresh; \
#     fi

# EXPOSE 8080