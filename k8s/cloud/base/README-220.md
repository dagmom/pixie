kubectl create secret generic -n "plc" \
  cloud-auth-secrets \
  --from-literal=jwt-signing-key="$(LANG=C; < /dev/urandom tr -dc 'a-zA-Z0-9' | fold -w 64 | head -n 1)"

kubectl create secret generic -n "plc" \
  pl-db-secrets \
  --from-literal=PL_POSTGRES_USERNAME="pl" \
  --from-literal=PL_POSTGRES_PASSWORD="pl" \
  --from-literal=database-key="$(LANG=C; < /dev/urandom tr -dc 'a-zA-Z0-9#$%&().' | fold -w 24 | head -n 1)"

kubectl create secret generic -n "plc" \
  cloud-session-secrets \
  --from-literal=session-key="$(LANG=C; < /dev/urandom tr -dc 'a-zA-Z0-9' | fold -w 24 | head -n 1)"

kubectl create secret generic -n "plc" \
  pl-hydra-secrets \
  --from-literal=SECRETS_SYSTEM="$(LANG=C; < /dev/urandom tr -dc 'a-zA-Z0-9' | fold -w 64 | head -n 1)" \
  --from-literal=OIDC_SUBJECT_IDENTIFIERS_PAIRWISE_SALT="$(LANG=C; < /dev/urandom tr -dc 'a-zA-Z0-9' | fold -w 64 | head -n 1)" \
  --from-literal=CLIENT_SECRET="$(LANG=C; < /dev/urandom tr -dc 'a-zA-Z0-9' | fold -w 64 | head -n 1)"  


kubectl create secret generic -n "plc" \
  service-tls-certs \
  --from-file=ca.crt=./ca.crt \
  --from-file=client.crt=./client.crt \
  --from-file=client.key=./client.key \
  --from-file=server.crt=./server.crt \
  --from-file=server.key=./server.key


PROXY_TLS_CERTS="/tmp/tmp.bz5YApLl4d"
PROXY_CERT_FILE="/tmp/tmp.bz5YApLl4d/server.crt"
PROXY_KEY_FILE="/tmp/tmp.bz5YApLl4d/server.key"

mkcert \
  -cert-file "/tmp/tmp.bz5YApLl4d/server.crt" \
  -key-file "/tmp/tmp.bz5YApLl4d/server.key" \
  dev.withpixie.dev "*.dev.withpixie.dev" localhost 127.0.0.1 ::1

kubectl create secret tls -n "plc" \
  cloud-proxy-tls-certs \
  --cert="/tmp/tmp.bz5YApLl4d/server.crt" \
  --key="/tmp/tmp.bz5YApLl4d/server.key"