echo "正常情况"
curl -X POST  \
    -F "title=TitleValue" \
    -F "content=ContentValue" \
    -F "preview=PreviewValue" \
    -F "likes=13271" \
    "localhost:13271/api/articles" 
echo "\n============\n"


echo "异常情况1: 缺失"
curl -X POST  \
    -F "Title=TitleValue" \
    -F "Content=ContentValue" \
    -F "Preview=" \
    -F "Likes=13271" \
    "localhost:13271/api/articles" 
echo "\n============\n"
