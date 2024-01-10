FROM scratch
COPY mockpi /
ENTRYPOINT ["/mockpi"]
