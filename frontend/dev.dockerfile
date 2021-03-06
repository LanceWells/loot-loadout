# pull official base image
FROM node:18.4.0-alpine3.15 AS test

# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
COPY package.json ./
COPY yarn.lock ./
COPY tsconfig.json ./
COPY README.md ./
RUN npm install --global yarn --force
RUN yarn

# add app
COPY src ./src
COPY public ./public

EXPOSE 3000

CMD ["yarn", "dev"]
