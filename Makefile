test:
	mysql -utestuser -ptestuser mdtd_bootcamp < table_create.sql
	export DATA_SOURCE_NAME="testuser:testuser@tcp(127.0.0.1:3306)/mdtd_bootcamp" ;\
	go test ./...
