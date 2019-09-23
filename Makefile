PROJECT_ID ?=

cloudrun:
	gcloud builds submit --tag gcr.io/$(PROJECT_ID)/ip
