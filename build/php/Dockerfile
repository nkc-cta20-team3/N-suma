FROM php:8.2.0-apache

# COPY --from=composer:latest /usr/bin/composer /usr/bin/composer

WORKDIR /var/www/

RUN apt-get update && apt-get install -y libonig-dev \
  && docker-php-ext-install pdo_mysql mysqli 

COPY ./php.ini /usr/local/etc/php/php.ini