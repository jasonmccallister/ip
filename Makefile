PROJECT_ID ?= ip-api-253801

cloudrun:
	gcloud builds submit --tag gcr.io/$(PROJECT_ID)/ip
