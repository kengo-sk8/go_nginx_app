FROM nginx:latest

RUN apt-get update && \ 
    apt-get upgrade -y && \ 
    apt-get install -y \ 
    curl && \ 
    rm -rf /var/lib/apt/lists/*

# default.confを削除
RUN rm -f /etc/nginx/conf.d/default.conf

# Nginxの設定ファイルをコピー
COPY ./conf/web1.conf /etc/nginx/conf.d/web1.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
