#! /bin/bash

output=$(opa test rego/ -v)
if [ $? -eq 0 ]
then
  printf "Tests PASSED\n$output\n"
else
  printf "Tests FAILED\n$output\n"
  exit $?
fi
tar -C data -zcvf bundles/data/bundle.tar.gz .
echo "bundled data - bundles/data/bundle.tar.gz"
tar --exclude "*_test.rego" -C rego -zcvf bundles/rego/bundle.tar.gz .
echo "bundled rego - bundles/rego/bundle.tar.gz"