# Start from node base image
FROM node:12-alpine

# Set the current working directory inside the container
WORKDIR /usr/src/app

# Copy package.json, yarn.lock files and download deps
COPY package.json yarn.lock ./
RUN yarn global add @angular/cli
RUN yarn

# Copy sources to the working directory
COPY . .

# Set the node environment
ARG node_env=development
ENV NODE_ENV $node_env

# Run the app
CMD ["ng", "run", "all:serve"]
