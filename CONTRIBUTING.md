## remake ssl/tls certificate

```sh
docker run --rm --name temp_certbot \
    -v ./nginx/certbot/conf:/etc/letsencrypt \
    -v ./nginx/certbot/www:/tmp/letsencrypt \
    -v ./nginx/certbot/log:/var/log \
    certbot/certbot:latest \
    certonly --webroot --agree-tos --renew-by-default \
    --preferred-challenges http-01 --server https://acme-v02.api.letsencrypt.org/directory \
    --text --email  hotorcelexo@gmail.com \
    -w /tmp/letsencrypt -d ctf01d.ru
```
