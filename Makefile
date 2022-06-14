proto-gen:
	cd api && \
	buf generate

manifests-gen:
	kubectl kustomize ./backend/publicapi/config/overlays/prod/ -o manifests/publicapi.yaml && \
	kubectl kustomize ./backend/roomsocket/config/overlays/prod/ -o manifests/roomsocket.yaml
