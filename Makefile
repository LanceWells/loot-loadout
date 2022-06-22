gen-proto:
	cd api && \
	buf generate

gen-manifests:
	kubectl kustomize ./backend/publicapi/config/overlays/prod/ -o manifests/publicapi.yaml && \
	kubectl kustomize ./backend/roomsocket/config/overlays/prod/ -o manifests/roomsocket.yaml && \
	kubectl kustomize ./frontend/config/overlays/prod/ -o manifests/frontend.yaml

get-kubeapps-secret:
	kubectl get --namespace default secret kubeapps-operator-token -o jsonpath='{.data.token}' -o go-template='{{.data.token | base64decode}}' && echo

storybook:
	cd frontend && \
	yarn storybook
