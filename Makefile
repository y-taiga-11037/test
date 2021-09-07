test:
	mysql -utestuser -ptestuser mdtd_bootcamp < table_create.sql
<<<<<<< 1ad9422131c064f27bf192165d4eab70d1a8f1d6
<<<<<<< c5c52fe5feaf62093f0cb5207c30a2d3b4a347e3
	export DATA_SOURCE_NAME="testuser:testuser@tcp(127.0.0.1:3306)/mdtd_bootcamp" ;\
=======
	DATA_SOURCE_NAME=testuser:testuser@tcp(127.0.0.1:3306)/mdtd_bootcamp
>>>>>>> make api
=======
	export DATA_SOURCE_NAME="testuser:testuser@tcp(127.0.0.1:3306)/mdtd_bootcamp"  ;\
>>>>>>> fixup! make api
	go test ./...
