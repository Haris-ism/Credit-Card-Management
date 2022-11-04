FROM node:10.16.3-alpine
RUN npm install -g --unsafe-perm node-red@0.20.8
WORKDIR /app
COPY package.json .
RUN npm install
COPY . .
EXPOSE 1880
# CMD ["node-red", "-u", "./"]
CMD ["node-red", "-s", "settings.js", "-u", "./", "flows_iot-server.json"]
