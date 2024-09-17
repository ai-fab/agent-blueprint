#!/bin/bash

# Set the base URL
BASE_URL="http://localhost:8090/api"

# Test client credentials
CLIENTS=(
    "test_client_1:test_secret_1"
    "test_client_2:test_secret_2"
)

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Function to make API calls
call_api() {
    local client_id=$1
    local client_secret=$2
    local method=$3
    local endpoint=$4
    local data=$5
    
    curl -s -X $method \
         -H "Content-Type: application/json" \
         -H "X-Client-ID: $client_id" \
         -H "X-Client-Secret: $client_secret" \
         -d "$data" \
         $BASE_URL$endpoint
}

for CLIENT in "${CLIENTS[@]}"; do
    IFS=':' read -r CLIENT_ID CLIENT_SECRET <<< "$CLIENT"
    
    echo -e "\n${GREEN}Testing client: $CLIENT_ID${NC}"

    # Test creating a project
    echo -e "${GREEN}Testing CREATE project${NC}"
    CREATE_RESPONSE=$(call_api $CLIENT_ID $CLIENT_SECRET POST "/projects" "{\"name\":\"Test Project for $CLIENT_ID\", \"status\":\"active\"}")
    echo $CREATE_RESPONSE
    PROJECT_ID=$(echo $CREATE_RESPONSE | jq -r '.id')

    # Test listing projects
    echo -e "\n${GREEN}Testing LIST projects${NC}"
    call_api $CLIENT_ID $CLIENT_SECRET GET "/projects"

    # Test getting project status
    echo -e "\n${GREEN}Testing GET project status${NC}"
    call_api $CLIENT_ID $CLIENT_SECRET GET "/projects/$PROJECT_ID/status"
done

echo -e "\n${GREEN}CRUD tests completed${NC}"
