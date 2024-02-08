FROM node:14.21 AS ANGULAR_BUILD
WORKDIR /app
COPY package*.json .
RUN npm install -g @angular/cli@8.3.12
COPY . .
RUN npm install  && ng build --prod

FROM nginx:alpine
COPY --from=ANGULAR_BUILD /app/dist/webapp /usr/share/nginx/html
EXPOSE 80 4200

CMD ["nginx", "-g", "daemon off;"]