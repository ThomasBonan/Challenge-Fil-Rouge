FROM ubuntu
RUN apt update
RUN apt install -y apache2
RUN apt install -y apache2-utils
RUN apt install -y libapache2-mod-php
RUN apt install -y php
RUN apt clean
COPY Site/index.php /var/www/html/index.php
COPY Site/lang.php /var/www/html/lang.php
COPY Site/lang/en.php /var/www/html/lang/en.php
COPY Site/lang/fr.php /var/www/html/lang/fr.php
COPY Site/lang/es.php /var/www/html/lang/es.php
COPY Site/flags/flag.txt /var/www/html/flags/flag.txt
EXPOSE 80
CMD ["apache2ctl", "-D", "FOREGROUND"]

