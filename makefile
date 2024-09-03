client:
	@go build -o bin/client client/main.go
	@./bin/client

game_producer:
	@go build -o bin/producer ./game_score_producer/.
	@./bin/producer

game_consumer:
	@go build -o bin/consumer ./game_score_consumer/.
	@./bin/consumer

cache:
	@go build -o bin/cache ./leaderboard_cache/.
	@./bin/cache

leaderboard_consumer:
	@go build -o bin/leaderboard ./leaderboard_cache/.
	@./bin/leaderboard