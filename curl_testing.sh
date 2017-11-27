#!/bin/bash
curl -H "Content-type: application/json" -X POST --data-binary '{"id": "2", "firstname": "Maria", "lastname": "Raboy"}' http://localhost:9876/people/2
sleep 1
curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:9876/people/2
sleep 1
curl  -X DELETE http://localhost:9876/people/2
echo "After delete"
sleep 1
curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:9876/people/2


curl -H "Content-type: application/json" -X POST --data-binary '{"id": "1", "firstname": "Nic", "lastname": "Raboy", "address": {"city": "Dublin", "state": "CA"}}' http://localhost:9876/people/101
sleep 1
curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:9876/people/101
curl -i -X DELETE http://localhost:9876/people/101
echo "After delete"
sleep 1
curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://localhost:9876/people/101
