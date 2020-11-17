FROM alpine

RUN apk add build-base && \
    mkdir -p /tmp/mecab && \
    wget https://github.com/shogo82148/mecab/releases/download/v0.996.5/mecab-0.996.5.tar.gz -O- | tar zxf - -C /tmp/mecab --strip-components 1 && \
    cd /tmp/mecab && \
    ./configure && \
    make && \
    make install

RUN apk add bash curl openssl sudo && \
    mkdir -p /tmp/neologd && \
    wget https://github.com/neologd/mecab-ipadic-neologd/archive/master.tar.gz -O- | tar zxf - -C /tmp/neologd --strip-components 1 && \
    /tmp/neologd/bin/install-mecab-ipadic-neologd -n -y -a && \
    echo "dicdir = /usr/local/lib/mecab/dic/mecab-ipadic-neologd" > /usr/local/etc/mecabrc

COPY . /mecab-grpc

RUN apk add go git && \
    cd /mecab-grpc && \
    CGO_LDFLAGS="$(mecab-config --libs)" \
    CGO_CFLAGS="-I$(mecab-config --inc-dir)" \
    go build -tags mecab -o /usr/local/bin/mecab-grpc

FROM alpine

RUN apk add --no-cache libstdc++

COPY --from=0 /usr/local/lib/libmecab.so* /usr/local/lib/
COPY --from=0 /usr/local/lib/mecab /usr/local/lib/mecab
COPY --from=0 /usr/local/etc/mecabrc /usr/local/etc/mecabrc
COPY --from=0 /usr/local/bin/mecab-grpc /usr/local/bin/mecab-grpc

EXPOSE 9000

ENTRYPOINT ["mecab-grpc"]
CMD ["serve", "--listen", ":9000"]
