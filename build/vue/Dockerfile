FROM node:18-alpine

WORKDIR /app/N-suma

COPY ./vite.sh /vite.sh
RUN chmod +x /vite.sh
RUN nkf -wLu --overwrite /vite.sh
CMD [ "sh","/vite.sh" ]
