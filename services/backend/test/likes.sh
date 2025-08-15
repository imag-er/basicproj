
echo "点赞"
curl -X POST  \
    "localhost:13271/api/likes/1" 

echo "\n"

echo "获取所有likes信息"
curl -X GET  \
    "localhost:13271/api/likes/1" 


