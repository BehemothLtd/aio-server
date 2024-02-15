MIGRATIONS_FOLDER = $(PWD)/database/migrations
DB_USERNAME=root
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=snippets
DB_PASSWORD=
DB_URL=mysql://$(DB_USERNAME):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_DATABASE)?parseTime=true&loc=Local

db.migrate:
	migrate -path $(MIGRATIONS_FOLDER) -database "${DB_URL}" up

db.migrate.rollback:
	migrate -path $(MIGRATIONS_FOLDER) -database "${DB_URL}" down