#!/bin/bash

# Set the base URL
BASE_URL="http://localhost:8090/api"

# Test client credentials
CLIENT_ID="test_client"
CLIENT_SECRET="test_secret"

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Function to make API calls
call_api() {
    local method=$1
    local endpoint=$2
    local data=$3
    
    curl -s -X $method \
         -H "Content-Type: application/json" \
         -H "X-Client-ID: $CLIENT_ID" \
         -H "X-Client-Secret: $CLIENT_SECRET" \
         -d "$data" \
         $BASE_URL$endpoint
}

# Test creating a project
echo -e "${GREEN}Testing CREATE project${NC}"
CREATE_RESPONSE=$(call_api POST "/projects" '{"name":"Test Project", "status":"active"}')
echo $CREATE_RESPONSE
PROJECT_ID=$(echo $CREATE_RESPONSE | jq -r '.id')

# Test listing projects
echo -e "\n${GREEN}Testing LIST projects${NC}"
call_api GET "/projects"

# Test getting project status
echo -e "\n${GREEN}Testing GET project status${NC}"
call_api GET "/projects/$PROJECT_ID/status"

echo -e "\n${GREEN}CRUD tests completed${NC}"
