# Token generator for Docker Registry v2

This is a token generator program for Docker Registry v2.


## Usage

```
$ ./docker-registry-token-genarator -h
Usage of ./docker-registry-token-genarator:
  -issuer string
        Issuer string for token (default "distribution-token-server")
  -key string
        Private key file
  -scope string
        scope
  -service string
        service
  -username string
        username


$ ./docker-registry-token-genarator -key ./cert.key -service registry.docker.io \
  -scope repository:samalba/my-app:pull,push -username mizzy
eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp3ayI6eyJlIjoiQVFBQiIsImtpZCI6IlZDSlM6T1pYUDpOSTRWOlhDUlI6RlBPRDpGQVRLOjZaMlc6MkZCTzo0SVlYOlA0TE06WTM1MzpNM1o3Iiwia3R5IjoiUlNBIiwibiI6InV6aDNoQXR6T3JySWotS2VMcjlNUDFnVkxWYlVURndjRFJqWkFCTFNiNGxyaHBhazVCRW9oWEVES1dYc0VtbHI5clB2eDF5aDMtRWhWZTdUTnhQNkdSUXBrczRSSUxoQmZZclktLU5CQl9UcGdGOEpXVWF1aFM5T2ZzeHd6MFNGNHA4N2o4azZRTngtR1NxUjhMOTRCaG9CS1R5aGZ2UXJCZjdDMWFUbFp4eDZFanpoVkZkei1BNk5Jd3YyR3ExYlNvNjUtNDhMRVlqNzZXVkVvVnZJQkY5MFJaMnNrOW5UMWx1UGUzTGxGbFRWV2F3ZHlMQUJ5YnE0c1VnZUJGejJHaVZnTWZEd0tnclpDRTloQkJhRFNDUEVkeTRfS0RlZ2t6S0YyTlVLMXRnVUI4SVVGcURzVDdWX3ZYYnhtUVI2ZEt2cVFqeHU1OWUzbFVYbTg5WUhXQWRmeEVleDQwUXZLa0MyUHk1RU5RRHpWZlF6RXdweWxIUnVJUUJFM3Q2ODg3LXVvbDVtTFNENnZoMWl2MXZVTVlZYTdPUVc4N2NQWEEyZDJjM3NZek5TS2VQOV9YLUdNTGdudThFOVZfZU1FUTFKU3MzZzB0aHJ3U2RvSUFkM01CNWRDbnJWc1VPX2VucW9zaWZOV3hIN3lGLXR3Qjdpc29yeTZjVDlucU9IY3lNTHhuazJZX3FvenJ3VlJHWk03ZTdIR3dzMG9WemR5cHFzMEJaME1qVGhzc1JkTkRScVJiQWd6cVRlTDdrM0kzbDBQM0hWeG04QnZvTUlubkpYY2JZMHpFZlFDSEZvX0xUNGZhWnNBTTRNZkIxNGFCYi02U3ZFSWxFVkdGcVd5ekZzLU80XzFIa1dOQzlRSlQ0Rmc4cFpFSE1WVmQ5azBnUXBSaGlRS0RrIn19.eyJpc3MiOiJkaXN0cmlidXRpb24tdG9rZW4tc2VydmVyIiwic3ViIjoibWl6enkiLCJhdWQiOiJyZWdpc3RyeS5kb2NrZXIuaW8iLCJleHAiOjE0NjQ3NTQ2NTksIm5iZiI6MTQ2NDc1NDM1OSwiaWF0IjoxNDY0NzU0MzU5LCJqdGkiOiI3ZXFxZGdWMGZpWUNWSUZyLWVPciIsImFjY2VzcyI6W3sidHlwZSI6InJlcG9zaXRvcnkiLCJuYW1lIjoic2FtYWxiYS9teS1hcHAiLCJhY3Rpb25zIjpbInB1bGwiLCJwdXNoIl19XX0.Oapc_ZVr4mludb0prSmSitjrtHqZAgTpWwPhoFv6rzh1GMB5_hTvx4EhO1GVOQj20P-2CEAze_2DbevRsqqEtK7rfY3v1rmzZCf0eIrl6-j5PFDimH8JE4CTVbnNbbVh57SgVg05WI70HcdZ0b14ccrCOJ_byfE6LqrzOmZFkU4FgGkPlCXp7aTf7_jgyN3B7-Eq8Z7CJ24xYLRcGfFGFRLqLXNaayksoKMN1XYvwGg1oxWVoORQ8TjjrVnG_0hLqXYpuB32_gD6yr5eTUakx7PTZ2pzu0lr1qIva0KI1bQrbuRyN4PIrAMrmlRcjKazwwk9GDnKv5Hgm6QO4CA6Uo6ZpVBJmucCaZPrDWOQKDH0yzQGqttC96IBSO4d-_v-tpFe42-O8thkCMANtbIl8B4W90hbIPJtOwlH95MHHWRIM5RXeFSgZm6UASyFN49KjJ6jIlkzroLoORD3mA-je2JEZF5TSW974d9MT2FcC004Po3hdJiMv2GjaUGdXnS2qJkzpaKT7mmDaQ3PrbVbb9PoGAm_5FqHzj_VZA9kT6zHTOSetwocSphKSZ0F63Dd0QhSVrCbfYJo2pbmKPfKCzBBXAD5RUHvwkwT_d1TZTBsIQJVyaszVtTmFhpHtdozoqADWR86-bDL-5ox3A1DEa8MqhYyScH06mN7JRIvcsA
```

----

## How to use

I show an example how to use this token generator.


This is a simple token server PHP progoram. This sample does not have an authentication logic.

```php
<?php

header('Content-Type: application/json');

$user    = $_GET['account'];
$service = $_GET['service'];

$command = "./docker-registry-token-genarator -key ./cert.key -service $service -username $user";

if (isset($_GET['scope'])) {
    $command .= ' -scope ' .$_GET['scope'];
}

$token = `$command`;

echo "{ \"token\": \"$token\" }";
```

`cert.key` is a private key to sign tokens.

You can create a private key and a certificate by `create-certs.sh` of [gesellix/inject-docker-certs](https://github.com/gesellix/inject-docker-certs) easily.

Run this with built-in web server.

```
$ php -S 192.168.1.17:5050 token.php
```

Run docker registry

```
$ docker run -p 5000:5000 --restart=always --name registry \
  -v `pwd`/certs:/certs \
  -e REGISTRY_AUTH=token \
  -e REGISTRY_AUTH_TOKEN_REALM=http://192.168.1.17:5050/ \
  -e REGISTRY_AUTH_TOKEN_SERVICE="registry.docker.io" \
  -e REGISTRY_AUTH_TOKEN_ISSUER="distribution-token-server" \
  -e REGISTRY_AUTH_TOKEN_ROOTCERTBUNDLE=/certs/cert.cert \
  registry:2
```

`cert.cert` must include a public key of `cert.key` used by `docker-registry-token-generator` command.

Run `docker login`

```
$ docker login localhost:5000
Username: test
Password:
Login Succeeded
```

The sample token server does not have authentication logic, so any username and password would be accepted.

Pull an image from the official registry.

```
$ docker pull ubuntu
```

Push the image to the authenticated registry.

```
$ docker push localhost:5000/ubuntu
The push refers to a repository [localhost:5000/ubuntu]
5f70bf18a086: Layer already exists
a3b5c80a4eba: Layer already exists
7f18b442972b: Layer already exists
3ce512daaf78: Layer already exists
7aae4540b42d: Layer already exists
latest: digest: sha256:15e2ba0b5e9d02f62a16332bf714732202d14b5930be31ec84421ee90ecff822 size: 1334
```

----


## See also

* https://docs.docker.com/registry/spec/auth/jwt/
* https://github.com/docker/distribution/tree/master/contrib/token-server
