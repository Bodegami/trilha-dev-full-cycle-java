FROM php:7.4-cli

WORKDIR /var/www

# Atualiza as libs do container, instala a lib zip e a extensao zip para php
RUN apt-get update && \ 
    apt-get install libzip-dev -y && \
    docker-php-ext-install zip

# Instala o gerenciador de pacotes "composer" do php
RUN php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');" && \
    php composer-setup.php && \
    php -r "unlink('composer-setup.php');"

# Instala o framework Laravel e cria um projeto 
RUN php composer.phar create-project --prefer-dist laravel/laravel laravel

ENTRYPOINT [ "php","laravel/artisan","serve" ]

CMD [ "--host=0.0.0.0" ]