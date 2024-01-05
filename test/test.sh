#! /bin/bash
echo ""
echo "---------------------------"
echo "testing HelloWorld template"
echo "---------------------------"
echo ""

curl -X GET -H "Content-Type: application/json" -d @hello-world-request.json localhost:8080/generateDocument

echo ""
echo "---------------------------"
echo "testing nda template"
echo "---------------------------"
echo ""

curl -X GET -H "Content-Type: application/json" -d @nda-request.json localhost:8080/generateDocument
