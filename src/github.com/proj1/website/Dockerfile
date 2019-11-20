FROM nginx
RUN apt-get update -y
RUN apt-get install nginx apt-utils net-tools -y
WORKDIR /usr/share/nginx/html
COPY ./* /usr/share/nginx/html/
#CMD ["service", "nginx start"]
EXPOSE 80
