FROM node as compile

WORKDIR /app
COPY package.json .
RUN npm i
COPY . .
RUN npm run build

from nginx as prod

COPY front.conf /etc/nginx/templates/default.conf.template
COPY --from=compile /app/dist /code
