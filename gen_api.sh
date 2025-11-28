
#!/bin/bash

rm -fr tmp

mkdir -p tmp/apimod

rm -fr model/apimod

openapi-generator-cli generate \
    -i openapi.json \
    -g go-gin-server \
    -o ./tmp \
    --additional-properties=packageName=apimod \
    --global-property models,modelDocs=false \
    --skip-validate-spec

mv tmp/go model/apimod
rm -fr tmp