#!/bin/bash

for i in {1..1000}
do
  curl -s -X POST http://localhost:6969/user/save \
    -H "Content-Type: application/json" \
    -d '{
      "username": "ismarianto",
      "password": "rahasia123",
      "email": "ismarianto@example.com",
      "created_at": "2025-09-19T21:45:00Z",
      "updated_at": "2025-09-19T22:00:00Z"
    }'

  echo " <- Request ke-$i selesai"
done

