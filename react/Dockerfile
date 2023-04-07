FROM node

WORKDIR /app/react

COPY package*.json /app/react/

RUN npm install

COPY . .

RUN npm run build

EXPOSE 3000

ENTRYPOINT [ "npm", "start" ]

