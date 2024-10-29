.PHONY apply:
apply:
	# @gcloud auth login --brief
	@kubectl apply -f namespace.yml
	@kubesec decrypt secret.yml -o decrypted-secret.yml
	@kubectl apply -f decrypted-secret.yml
	@kubectl apply -f deployment.yml
	@kubectl apply -f tunnel_api.yml