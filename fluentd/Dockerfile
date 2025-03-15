# fluentd/Dockerfile

FROM fluent/fluentd:edge-debian
USER root
RUN gem install elasticsearch -v 7.13.1
RUN gem install fluent-plugin-elasticsearch -v 4.0.0
RUN gem install fluent-plugin-rewrite-tag-filter
RUN gem install fluent-plugin-multi-format-parser
USER fluent