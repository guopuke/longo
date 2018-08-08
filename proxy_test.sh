#!/bin/bash

for n in $(seq 1 1 10)
do
    nohup curl -XGET -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MzM0NDU4MjksImlkIjo3LCJuYmYiOjE1MzM0NDU4MjksInVzZXJuYW1lIjoianNvbkhlbGxvIn0.JODlPjYRLbiyCz1e9SiXxMtcT4p4ZCsrBIru0IvkxJE" http://apiserver.com/v1/user &>/dev/null
done