#!/bin/bash
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

set -e

rm -rf "$SCRIPT_DIR/build/certs" && mkdir -p "$SCRIPT_DIR/build/certs"
cd "$SCRIPT_DIR/build/certs"

SUBJECT_PREFIX="/C=CA/ST=ONTARIO/L=TORONTO/O=ORG/OU=UNIT"
# Root CA
openssl genrsa -out root-ca-key.pem 2048
openssl req -new -x509 -sha256 -key root-ca-key.pem -subj "$SUBJECT/CN=opensearch-cluster-root" -out root-ca.pem -days 730
# Node cert 1
openssl genrsa -out node1-key-temp.pem 2048
openssl pkcs8 -inform PEM -outform PEM -in node1-key-temp.pem -topk8 -nocrypt -v1 PBE-SHA1-3DES -out node1-key.pem
openssl req -new -key node1-key.pem -subj "$SUBJECT_PREFIX/CN=opensearch-cluster-node1" -out node1.csr
echo "subjectAltName=DNS:opensearch-cluster-node1" > node1.ext
openssl x509 -req -in node1.csr -CA root-ca.pem -CAkey root-ca-key.pem -CAcreateserial -sha256 -out node1.pem -days 730 -extfile node1.ext
# Client cert
openssl genrsa -out client-key-temp.pem 2048
openssl pkcs8 -inform PEM -outform PEM -in client-key-temp.pem -topk8 -nocrypt -v1 PBE-SHA1-3DES -out client-key.pem
openssl req -new -key client-key.pem -subj "$SUBJECT_PREFIX/CN=opensearch-cluster-dashboard" -out client.csr
echo "subjectAltName=DNS:opensearch-cluster-dashboard" > client.ext
openssl x509 -req -in client.csr -CA root-ca.pem -CAkey root-ca-key.pem -CAcreateserial -sha256 -out client.pem -days 730 -extfile client.ext
# Cleanup
rm node1-key-temp.pem
rm node1.csr
rm node1.ext
rm client-key-temp.pem
rm client.csr
rm client.ext