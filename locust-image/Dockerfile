
# Start with a base Python 2.7 alpine image
FROM python:2.7-alpine

COPY requirements.txt requirements.txt
RUN apk --no-cache add --virtual=.build-dep \
      build-base \
    && apk --no-cache add bash libzmq \
    && pip install -r requirements.txt \
    && apk del .build-dep

COPY entrypoint.sh entrypoint.sh

# Set script to be executable
RUN chmod 755 /entrypoint.sh

# Expose the required Locust ports
EXPOSE 5557 5558 8089

# Start Locust using LOCUS_OPTS environment variable
ENTRYPOINT ["/entrypoint.sh"]