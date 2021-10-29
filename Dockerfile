# Download latest golang image
FROM golang:latest
# Create a directory for the app
RUN mkdir /app
# Copy all files from current directory to working directory
COPY . /app
# Set working directory
WORKDIR /app

RUN cd cmd && go build -o Nourybot

CMD ["/app/cmd/Nourybot"]