
echo "正常情况"
curl -X POST  \
    -H "Authorization: Bearer $1" \
    -H "Content-Type: application/json" \
    "localhost:13271/api/exchange-rates" 
echo "\n============\n"


echo "异常情况"
curl -X POST  \
    "localhost:13271/api/exchange-rates" 
echo "\n============\n"