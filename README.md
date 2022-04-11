# netcheck
Script envía un email si el servidor está caído Protocolo TCP

## usage

```
1. configurar archivo .env
SERVER=
PORT=
FROM=
TO=
SUBJECT=
EMAIL=
PASSWORD=
```


```
2.
docker-compose up
```

## examples

```
 2022/04/11 14:22:11 Starting tcp port check: 192.168.102.214:22
2022/04/11 14:22:11 Connection success to "192.168.102.214:22"
2022/04/11 14:22:31 Connection success to "192.168.102.214:22"
2022/04/11 14:22:51 Connection success to "192.168.102.214:22"
2022/04/11 14:23:11 Connection success to "192.168.102.214:22"
2022/04/11 14:23:31 Connection success to "192.168.102.214:22"
2022/04/11 14:23:51 Connection success to "192.168.102.214:22"
2022/04/11 14:24:12 Connection success to "192.168.102.214:22"
2022/04/11 14:24:32 Connection success to "192.168.102.214:22"
2022/04/11 14:24:52 Connection success to "192.168.102.214:22"
2022/04/11 14:25:12 Connection success to "192.168.102.214:22"
```