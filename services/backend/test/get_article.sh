
echo "正常情况:all"
curl -X GET  \
    "localhost:13271/api/articles" 
echo "\n============\n"

echo "正常情况:byid"
curl -X GET  \
    "localhost:13271/api/articles/1" 
echo "\n============\n"

