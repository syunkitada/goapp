openssl genrsa -out server.key 2048
openssl ecparam -genkey -name secp384r1 -out server.key

cat << EOS > openssl.cnf
[req]
req_extensions = v3_req
distinguished_name = req_distinguished_name
[req_distinguished_name]
[v3_req]
basicConstraints = CA:FALSE
keyUsage = nonRepudiation, digitalSignature, keyEncipherment
subjectAltName = @alt_names
[alt_names]
DNS.1 = localhost
EOS

openssl req -new -x509 -sha256 -key server.key -out server.pem -subj "/CN=localhost" -days 3650 -config openssl.cnf
