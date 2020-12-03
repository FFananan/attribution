# Configration

## key manager

| name | description | default |
| --- | --- | --- |
| key_manager_name | key manager storage type: `redis` / `hdfs` | `redis` |
| default_redis_config | redis config | `{"address":["127.0.0.1:6379"],"is_cluster":1}` |
| hdfs_address | hdfs config | `""` |
| hdfs_user | hdfs config | `""` |

## impression data storage

| name | description | default |
| --- | --- | --- |
| imp_kv_type | impression data storage type: `LEVELDB` / `REDIS` | `LEVELDB` |
| imp_kv_address | impression data storage address. it is *leveldb*'s path by default | `/data/db` |
| imp_kv_password | impression data storage password. it is no password for *leveldb* by default  | `""` |

## metadata storage

| name | description | default |
| --- | --- | --- |
| store_type | metadata storage type: `SQLITE` / `REDIS` | `SQLITE` |
| store_option | metadata storage config | `{"dsn":"/data/sqlite.db"}` |

## server config

| name | description | default |
| --- | --- | --- |
| serverAddress | server address and port | `:80` |
| metric_server_address | metric for prometheus address and port | `:8080` |
| ams_encrypt_url | ams crypto server's encrypt api address | `http://localhost:9010/encrypt` |
| ams_decrypt_url | ams crypto server's decrypt api address | `http://localhost:9010/decrypt` |
| v | glog's v | `100` |