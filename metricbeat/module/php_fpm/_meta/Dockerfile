FROM richarvey/nginx-php-fpm:php71

RUN echo "pm.status_path = /status" >> /usr/local/etc/php-fpm.d/www.conf
ADD ./php-fpm.conf /etc/nginx/sites-enabled

EXPOSE 81
