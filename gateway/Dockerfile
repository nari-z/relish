FROM alpine:3.12.0

COPY ./gateway .
RUN chmod u+x gateway

EXPOSE 1111
CMD ["./gateway"]