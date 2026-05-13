# workout-management-service

Service that manages the creation and querying of workouts. Based off of github actions workshop with practical examples https://samirs-organization-6.gitbook.io/github-actions-by-example

## Running the service locally

1. Start local running dynamodb

```bash
# must be running docker to run this command
docker-compose up
```

2. Build the service

```bash
go build -o workout-management-service
```

3. Run the service

```bash
./workout-management-service
```

4. Troubleshooting tip for running create-table.sh
   Ensure local machine has awscli installed

```brew install awscli
aws configure
```

_This creates ~/.aws/credentials on your machine. DynamoDB Local doesn't validate these, but AWS CLI requires them._
When prompted, enter:

AWS Access Key ID: []
AWS Secret Access Key: []
Default region name: us-west-2
Default output format: json

## Making requests to the service locally

1. Create a workout

```bash
curl -X POST http://localhost:1323/create \
-H "Content-Type: application/json" \
-d '{
  "owner": "samir@gmail.com",
  "name": "Run The Interval",
  "category": "running",
  "equipment": {
    "name": "Running Equipment",
    "description": "If indoors we need a threadmill, if outdoors a good place to run fast for 3 min without interruption"
  },
  "exercises": [
    {
      "name": "Warmup",
      "description": "20min jog"
    },
    {
      "name": "3 min by 5 interval",
      "description": "3min x5 at 5k pace with 1min jog"
    },
    {
      "name": "Cooldown",
      "description": "20min cool down"
    }
  ]
}'

```

2. Get a workout

```bash
curl -X POST http://localhost:1323/get \
-H "Content-Type: application/json" \
-d '{
  "owner": "Samir",
  "name": "Run The Interval"
}'
```
