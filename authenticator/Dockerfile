FROM alpine:3.12.0

COPY ./authenticator .
RUN chmod u+x authenticator

EXPOSE 9999
CMD ["./authenticator"]