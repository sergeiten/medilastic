#!/usr/bin/env bash

# docker build -t elasticsearch-medicalpie .

# docker network create elastic

# docker stop elasticsearch && docker rm elasticsearch
# docker run -d --network=elastic -p 9200:9200 -v ~/data:/usr/share/elasticsearch/data --name=elasticsearch elasticsearch-medicalpie

docker stop kibana && docker rm kibana
docker run -d --network=medilastic_default -p 5601:5601 --name=kibana docker.elastic.co/kibana/kibana-oss:6.1.1

#docker run -p 5601:5601 --network=elastic -e ELASTICSEARCH_URL=http://localhost:9200 -e SERVER_NAME=kibana -e SERVER_HOST=localhost --name kibana docker.elastic.co/kibana/kibana-oss:6.1.1

#curl -X POST 'http://localhost:9200/_analyze' -H'Content-Type: application/json' -d '{"analyzer": "korean", "text": "한국어를 처리하는 예시입니닼ㅋㅋ"}'

#curl -X POST 'http://localhost:9200/_analyze' -H'Content-Type: application/json' -d '{"analyzer": "openkoreantext-analyzer", "text": "한국어를 처리하는 예시입니닼ㅋㅋ"}'
