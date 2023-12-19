# build environment
FROM node:18-alpine as build-stage
COPY ./vue/app/N-suma /app/N-suma
WORKDIR /app/N-suma
RUN yarn install
RUN yarn build

# production environment
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/N-suma/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
