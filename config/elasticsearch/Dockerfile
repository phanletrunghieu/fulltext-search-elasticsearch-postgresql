# https://github.com/elastic/elasticsearch-docker
FROM docker.elastic.co/elasticsearch/elasticsearch:7.1.1

# Add your elasticsearch plugins setup here
# Example: RUN elasticsearch-plugin install analysis-icu
USER root
RUN usermod -u 1100 elasticsearch
USER elasticsearch