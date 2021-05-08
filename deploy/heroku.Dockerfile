FROM redis:6.2-alpine

COPY server/static /static
COPY server/build/deezer-rank /deezer-rank

COPY run-heroku.sh /run-heroku.sh
RUN ["chmod", "+x", "/run-heroku.sh"]

ENTRYPOINT ["sh", "/run-heroku.sh"]
