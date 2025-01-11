FROM nginx:alpine

COPY ./docker/nginx/default.conf.template /etc/nginx/conf.d/default.conf.template
COPY ./docker/entrypoints/nginx.prod.sh /app/docker/entrypoints/nginx.prod.sh
COPY ./docker/nginx/certs /etc/nginx/certs

ENTRYPOINT ["sh", "/app/docker/entrypoints/nginx.prod.sh"]
CMD ["nginx", "-g", "daemon off;"]