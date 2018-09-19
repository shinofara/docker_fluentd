build:
	docker build -t a/a .

run:
	docker run -p 24284:24284 -e "GCP_PROJECT_ID=xxxx" -v ${PWD}/key:/var/secrets/google:ro a/a
