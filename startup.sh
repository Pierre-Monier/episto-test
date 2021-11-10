#!/bin/sh

if [[ ${APP_ENV} == production ]];
    then
        go build && app;
    else
        go run *.go;
    fi