PROJECT_ID ?= mccallisterio

cloudrun:
	gcloud builds submit --tag gcr.io/$(PROJECT_ID)/ip
