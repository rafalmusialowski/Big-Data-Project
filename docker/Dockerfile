# Start from the official Golang image
FROM golang:1.19

# Set the working directory
WORKDIR /app

# Copy the source files into the container
COPY src/mapreduce/mapper.go src/mapreduce/combiner.go src/mapreduce/reducer.go ./

# Compile the Mapper, Combiner, and Reducer
RUN go build -o /usr/local/bin/mapper mapper.go
RUN go build -o /usr/local/bin/combiner combiner.go
RUN go build -o /usr/local/bin/reducer reducer.go

# Default to bash shell for interactive use
CMD ["/bin/bash"]