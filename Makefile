gen-proto:
	cd api && \
	buf generate

gen-manifests:
	kubectl kustomize ./backend/publicapi/config/overlays/prod/ -o manifests/publicapi.yaml && \
	kubectl kustomize ./backend/roomsocket/config/overlays/prod/ -o manifests/roomsocket.yaml && \
	kubectl kustomize ./frontend/config/overlays/prod/ -o manifests/frontend.yaml
