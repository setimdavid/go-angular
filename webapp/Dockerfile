FROM node:12.11 AS ANGULAR_BUILD
WORKDIR /app
COPY package*.json .
RUN npm install -g @angular/cli@8.3.12
RUN npm install && ng build --prod
COPY . .

FROM alpine:3.10
EXPOSE 4200
WORKDIR /app
COPY --from=ANGULAR_BUILD /app/* ./app/
RUN ls
CMD ./server
