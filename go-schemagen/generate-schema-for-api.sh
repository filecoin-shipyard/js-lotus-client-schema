#! /bin/bash

__dirname="$(CDPATH= cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
output="${__dirname}/../mainnet/gen.js"
echo "Working ..."

rm ${output}

for api in Common FullNode StorageMiner Gateway Wallet Worker; do
  echo -n "module.exports.${api}Methods = " >> ${output}
  go run . $api >> ${output}
  echo -e "\n" >> ${output}
  echo " ... ${api}"
done

echo "Wrote generated API to ${output}, use 'standard --fix' to reformat"