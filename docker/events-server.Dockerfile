FROM default AS builder

WORKDIR /events-server

RUN chmod ugo+x .bin/events-server

CMD [".bin/events-server"]