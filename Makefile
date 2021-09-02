test:
	# DBテストで使用するもの
	mysql -utestuser -ptestuser mdtd_bootcamp < table_create.sql
	go test ./...
