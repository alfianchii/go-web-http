FROM redis:7.4.1

RUN apt-get update && apt-get install -y gettext-base && apt-get clean

COPY ./docker/entrypoints/generate-acl.sh /usr/local/bin/generate-acl.sh

ENTRYPOINT ["sh", "/usr/local/bin/generate-acl.sh"]