
echo "正常情况"
curl -X POST  \
    -H "Authorization: Bearer $1" \
    "localhost:13271/api/auth/logout" 
echo "\n============\n"

echo "异常情况"
curl -X POST  \
    "localhost:13271/api/auth/logout" 
echo "\n============\n"