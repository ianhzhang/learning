FROM node:8

# Create app directory
RUN mkdir -p /var/www/server

WORKDIR /var/www/server

# Install app dependencies
# A wildcard is used to ensure both package.json AND package-lock.json are copied
# where available (npm@5+)
COPY package*.json ./
COPY app.js ./
RUN npm install

EXPOSE 5000

CMD [ "npm", "start" ]
