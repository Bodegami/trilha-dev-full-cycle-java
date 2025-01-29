FROM golang:1.19

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

# Instala dependências e ferramentas específicas para o projeto
RUN go install github.com/golang/mock/mockgen@v1.5.0 && \
    go install github.com/spf13/cobra-cli@latest

# Atualiza os pacotes e instala SQLite3
RUN apt-get update && apt-get install -y sqlite3 && apt-get clean && rm -rf /var/lib/apt/lists/*

# Configuração de permissões e cache
RUN usermod -u 1000 www-data && \
    mkdir -p /var/www/.cache && \
    chown -R www-data:www-data /go && \
    chown -R www-data:www-data /var/www/.cache

# Define o usuário como www-data
USER www-data

# Define o comando padrão
CMD ["tail", "-f", "/dev/null"]
