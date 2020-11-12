FROM gunosy/neologd-for-mecab

RUN apk add go gcc

WORKDIR /mecab-grpc
COPY . .

RUN CGO_LDFLAGS="$(mecab-config --libs)" \
    CGO_CFLAGS="-I$(mecab-config --inc-dir)" \
    go build

FROM gunosy/neologd-for-mecab

COPY --from=0 /mecab-grpc/mecab-grpc /usr/local/bin/mecab-grpc

EXPOSE 9000

ENTRYPOINT ["mecab-grpc"]
CMD ["serve", "--config", "dicdir:/usr/lib/mecab/dic/neologd", "--listen", ":9000"]
