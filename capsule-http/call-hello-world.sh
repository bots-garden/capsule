curl -X POST http://localhost:8080 \
-H 'Content-Type: application/json; charset=utf-8' \
-d @- << EOF
{
    "name":"Bob Morane",
    "age":42
}
EOF
