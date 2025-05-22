# Configuration

> To enable redis globally access to the server

Path : /etc/redis/redis.conf

Update to :

bind 0.0.0.0
protected-mode no
port 6379

sudo systemctl restart redis