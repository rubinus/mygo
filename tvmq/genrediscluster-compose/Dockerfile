FROM redis
LABEL maintainer="rubinus.chu@gmail.com"

ENV REDIS_PORT 7000
EXPOSE $REDIS_PORT

ADD entrypoint.sh /usr/local/bin/
ADD redis.conf /usr/local/etc/

RUN chmod 755 /usr/local/bin/entrypoint.sh

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
CMD ["redis-server", "/usr/local/etc/redis.conf"]
