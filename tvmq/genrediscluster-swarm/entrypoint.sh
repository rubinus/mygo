#!/bin/sh

set -e

# first arg is `-f` or `--some-option`
# or first arg is `something.conf`
if [ "${1#-}" != "$1" ] || [ "${1%.conf}" != "$1" ]; then
    set -- redis-server "$@"
fi

# allow the container to be started with `--user`
if [ "$1" = 'redis-server' -a "$(id -u)" = '0' ]; then

    sed -i "s/REDIS_PORT/$REDIS_PORT/g" /usr/local/etc/redis.conf
    sed -i "s/ANNOUNCEIP/$ANNOUNCEIP/g" /usr/local/etc/redis.conf

    chown -R redis .
    exec gosu redis "$0" "$@"
fi

exec "$@"
