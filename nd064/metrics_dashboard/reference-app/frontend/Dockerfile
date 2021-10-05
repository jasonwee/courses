# Use the official Python image.
# https://hub.docker.com/_/python
FROM python:3.7-slim

# Copy local code to the container image.
ENV APP_HOME /app
ENV PROMETHEUS_MULTIPROC_DIR /tmp/prometheus-multiproc

RUN rm -rf /tmp/prometheus-multiproc
RUN mkdir -p /tmp/prometheus-multiproc

WORKDIR $APP_HOME
COPY . ./

# Install production dependencies.
RUN apt-get update && \
    pip install --upgrade pip && \
    pip install -r requirements.txt

CMD ["gunicorn", "--access-logfile", "-", "--error-logfile", "-", "-w", "4", "-b", "0.0.0.0:8080", "app:app"]
