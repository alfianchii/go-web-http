FROM nginx:alpine

COPY ./docker/nginx/default.conf.template /etc/nginx/conf.d/default.conf.template
COPY ./docker/nginx/option-ssl-nginx.conf /etc/nginx/conf.d/option-ssl-nginx.conf
COPY ./docker/entrypoints/nginx.prod.sh /app/docker/entrypoints/nginx.prod.sh
COPY ./docker/nginx/certs/fullchain.pem /etc/nginx/certs/fullchain.pem
COPY ./docker/nginx/certs/privkey.pem /etc/nginx/certs/privkey.pem

ENTRYPOINT ["sh", "/app/docker/entrypoints/nginx.prod.sh"]
CMD ["nginx", "-g", "daemon off;"]