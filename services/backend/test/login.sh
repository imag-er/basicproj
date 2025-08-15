
echo "正常情况"
curl -X POST  \
    -F "username=lsm" \
    -F "password=123" \
    "localhost:13271/api/auth/login" 
echo "\n============\n"

echo "异常情况1: 密码错误"
curl -X POST  \
    -F "username=lsm" \
    -F "password=13" \
    "localhost:13271/api/auth/login" 
echo "\n============\n"

echo "异常情况2: 用户名不存在"
curl -X POST  \
    -F "username=lm" \
    -F "password=123" \
    "localhost:13271/api/auth/login" 
echo "\n============\n"

echo "异常情况3: 用户名和密码都错误"
curl -X POST  \
    -F "username=l" \
    -F "password=13" \
    "localhost:13271/api/auth/login" 
echo "\n============\n"

echo "异常情况4: 用户名和密码都为空"
curl -X POST  \
    -F "username=" \
    -F "password=" \
    "localhost:13271/api/auth/login" 
echo "\n============\n"