# pull official base image
FROM node:18.4.0-alpine3.15 AS test

# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
COPY package.json ./
COPY package-lock.json ./
COPY tsconfig.json ./
COPY README.md ./
RUN npm install --silent
RUN npm install react-scripts@3.4.1 -g --silent

# add app
COPY src ./src
COPY public ./public

RUN ["ls", "-a"]
CMD ["npm", "run", "start"]