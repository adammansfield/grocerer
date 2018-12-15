FROM ubuntu:latest

ENV CONFIG="/config/ourgroceries-rest-api.json"
ENV PORT=8090

VOLUME /config
