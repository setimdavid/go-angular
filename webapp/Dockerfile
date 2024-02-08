FROM node:14.21 AS ANGULAR_BUILD
WORKDIR /app
COPY package*.json .
# RUN npm install -g @angular/cli@8.3.12
RUN npm install -g @angular/cli@latest
# RUN npm install  && ng build --prod
RUN npm install
COPY . .
RUN npm run build

FROM node:14.21.3-buster-slim
EXPOSE 4200
WORKDIR /app
COPY --from=ANGULAR_BUILD /app /app
CMD ["npm","start"]
