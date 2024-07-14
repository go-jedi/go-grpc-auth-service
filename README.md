# grpc-auth-service

"grpc-auth-service" is a microservice for authorized users using a clean architecture.

### architecture

![alt text](testdata/img.png)

### config

```yaml
logger:
  is_json: true
  add_source: false
  level: debug
  set_file: false
  file_name: logs/app.log
  max_size: 10
  max_backups: 3
  max_age: 7

jwt:
  secret_path: testdata/fileServer/jwt/.secret # path to file with secret key
  secret_hash_len: 30
  access_exp_at: 5 # minutes
  refresh_exp_at: 30 # days

postgres:
  host: localhost
  user: auth
  password: auth
  dbname: auth_db
  port: 54321
  sslmode: disable
  pool_max_conns: 10
  migrations_dir: file://migrations

grpcserver:
  host: 127.0.0.1
  port: 50053
```