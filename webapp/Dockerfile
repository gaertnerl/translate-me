# syntax=docker/dockerfile:1

FROM node:12.22.9 as frontend
WORKDIR /
COPY ./backend/src /backend/src
COPY ./frontend /frontend/

WORKDIR /frontend

# compile typescript
RUN npm ci
RUN npx webpack

# copy frontend files into static folder of webserver
ARG STATIC_DIR=/backend/src/satic
RUN mkdir $STATIC_DIR
RUN cp -r ./src/html $STATIC_DIR/html
RUN cp -r ./src/compiled_js $STATIC_DIR/compiled_js
RUN cp -r ./src/style $STATIC_DIR/style
RUN cp -r ./src/favicon.ico $STATIC_DIR/favicon.ico

FROM golang:1.19-alpine

# copy from previous stage
COPY --from=frontend /backend /backend
WORKDIR /backend/src

ARG REGISTER_SIM_ENDPOINT_KEY
ARG DB_HOST
ARG DB_PORT
ARG DB_PASSWORD
ARG DB_USERNAME
ARG DB_NAME

ENV REGISTER_SIM_ENDPOINT_KEY=${REGISTER_SIM_ENDPOINT_KEY}
ENV DB_HOST=${DB_HOST}
ENV DB_PORT=${DB_PORT}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_USERNAME=${DB_USERNAME}
ENV DB_NAME=${DB_NAME}

EXPOSE 8080

CMD ["go", "run", "."]
