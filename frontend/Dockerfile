FROM node:20-alpine

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .
RUN npm run build

# The built files will be in /app/dist
# Your backend will serve these files from the mounted volume