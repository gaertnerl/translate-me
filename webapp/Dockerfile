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
EXPOSE 8080

CMD ["go", "run", "."]
